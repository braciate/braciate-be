package entity

type Lkms struct {
	ID         string
	Name       string
	CategoryID string
	LogoFile   string
	Type       int
}

type LkmsType uint8

const (
	NoTypeLkms      LkmsType = 0
	TypePenalaran   LkmsType = 1
	TypeKerohanian  LkmsType = 2
	TypeOlahraga    LkmsType = 3
	TypeKesenian    LkmsType = 4
	TypeMinatKhusus LkmsType = 5
)

var LkmsTypeMap = map[LkmsType]string{
	TypePenalaran:   "Penalaran",
	TypeKerohanian:  "Kerohanian",
	TypeOlahraga:    "Olahraga",
	TypeKesenian:    "Kesenian",
	TypeMinatKhusus: "Minat Khusus",
}

func (l LkmsType) GetString() string {
	return LkmsTypeMap[l]
}

func (l LkmsType) GetInt() int {
	return int(l)
}
