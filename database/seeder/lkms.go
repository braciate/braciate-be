package seeder

import (
	"bytes"
	"fmt"
	"github.com/braciate/braciate-be/internal/entity"
	"github.com/braciate/braciate-be/internal/pkg/bcrypt"
	"github.com/braciate/braciate-be/internal/pkg/utils"
	"github.com/jmoiron/sqlx"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"net/textproto"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"sync"

	"github.com/tealeg/xlsx"
)

func LkmsBemSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/BEM"
	excelFilePath := "./users-BEM.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.NoTypeLkms)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsHimaSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/HIMA"
	excelFilePath := "./users-HIMA.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.NoTypeLkms)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsDpmSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/DPM"
	excelFilePath := "./users-DPM.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.NoTypeLkms)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsUkmKerohanianSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/UKM/kerohanian"
	excelFilePath := "./users-UKM-Kerohanian.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.TypeKerohanian)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsUkmKesenianSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/UKM/kesenian"
	excelFilePath := "./users-UKM-Kesenian.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.TypeKesenian)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsUkmMinatKhususSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/UKM/minat-khusus"
	excelFilePath := "./users-UKM-MinatKhusus.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.TypeMinatKhusus)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsUkmOlahragaSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/UKM/olahraga"
	excelFilePath := "./users-UKM-Olahraga.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.TypeOlahraga)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func LkmsUkmPenalaranSeeder(db *sqlx.DB) error {
	folderPath := "./ukm/UKM/penalaran"
	excelFilePath := "./users-UKM-Penalaran.xlsx"

	log.Printf("Starting to upload files from folder: %s", folderPath)
	err := uploadAllFilesInFolder(folderPath, db, excelFilePath, entity.TypePenalaran)
	if err != nil {
		log.Fatalf("Error uploading files: %v", err)
	}
	log.Println("Upload and saving process completed successfully!")

	return nil
}

func uploadToSupabase(filePath string) (string, error) {
	supabaseURL := os.Getenv("SUPABASE_URL_TEST")
	supabaseBucket := os.Getenv("SUPABASE_BUCKET_TEST")
	supabaseToken := os.Getenv("SUPABASE_TOKEN_TEST")

	file, err := os.Open(filePath)
	if err != nil {
		log.Printf("Error: failed to open file %s: %v", filePath, err)
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer file.Close()

	var buffer bytes.Buffer
	writer := multipart.NewWriter(&buffer)

	h := make(textproto.MIMEHeader)
	h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"; filename="%s"`, "file", file.Name()))
	h.Set("Content-Type", "application/png")

	part, err := writer.CreatePart(h)
	if err != nil {
		return "", fmt.Errorf("failed to create multipart part: %w", err)
	}

	if _, err := io.Copy(part, file); err != nil {
		return "", fmt.Errorf("failed to copy file to multipart: %w", err)
	}

	if err := writer.Close(); err != nil {
		return "", fmt.Errorf("failed to close writer: %w", err)
	}

	_, fileName := filepath.Split(filePath)

	url := supabaseURL + supabaseBucket + "/" + fileName

	request, err := http.NewRequest(http.MethodPost, url, &buffer)
	if err != nil {
		log.Printf("Error: failed to create request: %v", err)
		return "", fmt.Errorf("failed to create request: %w", err)
	}
	token := fmt.Sprintf("Bearer %s", supabaseToken)

	request.Header.Set("Authorization", token)
	request.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		log.Printf("Error: failed to send request: %v", err)
		return "", fmt.Errorf("failed to send request: %w", err)
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		log.Printf("Error: received non-200 response: %d", response.StatusCode)
		return "", fmt.Errorf("bad response status: %d", response.StatusCode)
	}

	fileURL := supabaseURL + supabaseBucket + "/" + fileName
	log.Printf("Successfully uploaded: %s", fileName)
	return fileURL, nil
}

func insertLkmsAndUser(db *sqlx.DB, v entity.Lkms, file *xlsx.File) error {
	nameWithoutExt := strings.TrimSuffix(v.Name, filepath.Ext(v.Name))

	randomPass, err := utils.GenerateRandomString(10)
	if err != nil {
		log.Printf("Error: failed to generate random password for user %s: %v", nameWithoutExt, err)
		return fmt.Errorf("failed to generate random password: %w", err)
	}

	hashPass, _ := bcrypt.HashPassword(randomPass)

	emailUsername := sanitizeNameForEmail(nameWithoutExt)
	email := emailUsername + "@braciate.ub.ac.id"

	_, err = db.Exec("INSERT INTO lkms (id, name, category_id, logo_file, type) VALUES ($1, $2, $3, $4, $5)",
		emailUsername, nameWithoutExt, v.CategoryID, v.LogoFile, v.Type)
	if err != nil {
		log.Printf("Error: failed to insert Lkms %s: %v", nameWithoutExt, err)
		return fmt.Errorf("failed to insert Lkms: %w", err)
	}

	_, err = db.Exec("INSERT INTO users (id, name, nim, email, faculty, study_program, role, password) VALUES ($1, $2, $3, $4, $5, $6, $7, $8)",
		emailUsername,
		nameWithoutExt,
		"230000000000001",
		email,
		"FILKOM",
		"Delegation",
		entity.UserRoleDelegation,
		hashPass)
	if err != nil {
		log.Printf("Error: failed to insert user for %s: %v", v.Name, err)
		return fmt.Errorf("failed to insert user: %w", err)
	}

	err = saveToExcel(file, nameWithoutExt, email, randomPass)
	if err != nil {
		log.Printf("Error: failed to save credentials to Excel for %s: %v", v.Name, err)
		return fmt.Errorf("failed to save credentials to Excel: %w", err)
	}

	log.Printf("Successfully inserted Lkms and user for: %s (Email: %s)", v.Name, email)
	return nil
}

func uploadAllFilesInFolder(folderPath string, db *sqlx.DB, excelFilePath string, lkmsType entity.LkmsType) error {
	wg := sync.WaitGroup{}

	file := xlsx.NewFile()
	sheet, err := file.AddSheet("Credentials")
	if err != nil {
		log.Printf("Error: failed to create Excel sheet: %v", err)
		return fmt.Errorf("failed to create Excel sheet: %w", err)
	}

	row := sheet.AddRow()
	row.AddCell().Value = "Name"
	row.AddCell().Value = "Email"
	row.AddCell().Value = "Password"

	err = filepath.Walk(folderPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("Error walking through folder: %v", err)
			return err
		}

		if !info.IsDir() {
			wg.Add(1)

			go func(path string, info os.FileInfo) {
				defer wg.Done()

				defer func() {
					if r := recover(); r != nil {
						log.Printf("Recovered from panic in goroutine for file %s: %v", path, r)
					}
				}()

				fileURL, err := uploadToSupabase(path)
				if err != nil {
					log.Printf("Error uploading file %s: %v", path, err)
					return
				}

				id, _ := utils.GenerateRandomString(24)

				lkms := entity.Lkms{
					ID:         id,
					Name:       info.Name(),
					CategoryID: "hima",
					LogoFile:   fileURL,
					Type:       lkmsType.GetInt(),
				}

				err = insertLkmsAndUser(db, lkms, file)
				if err != nil {
					log.Printf("Error inserting Lkms and user for file %s: %v", path, err)
					return
				}

				log.Printf("Successfully processed file: %s", path)

			}(path, info)
		}
		return nil
	})

	wg.Wait()

	if err != nil {
		log.Printf("Error walking through folder: %v", err)
		return fmt.Errorf("error walking through folder: %w", err)
	}

	err = file.Save(excelFilePath)
	if err != nil {
		log.Printf("Error saving Excel file: %v", err)
		return fmt.Errorf("failed to save Excel file: %w", err)
	}

	log.Printf("Credentials saved to: %s", excelFilePath)
	return nil
}

func saveToExcel(file *xlsx.File, name, email, password string) error {
	sheet := file.Sheets[0]
	row := sheet.AddRow()
	row.AddCell().Value = name
	row.AddCell().Value = email
	row.AddCell().Value = password

	log.Printf("Credentials saved for: %s (Email: %s)", name, email)
	return nil
}

func sanitizeNameForEmail(name string) string {
	sanitized := strings.ToLower(strings.ReplaceAll(name, " ", "_"))
	re := regexp.MustCompile("[^a-zA-Z0-9_]+")
	sanitized = re.ReplaceAllString(sanitized, "")
	return sanitized
}
