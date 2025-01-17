package votesRepository

const (
	queryCreateUserVotes = `
	INSERT INTO User_votes (id, user_id, lkm_id, nomination_id)
	VALUES (:id, :user_id, :lkm_id, :nomination_id)
	RETURNING id, user_id, lkm_id, nomination_id
	`

	queryGetUserVoteFromNominationID = `
	SELECT id, user_id AS UserID, nomination_id AS NominationID, lkm_id AS LkmID
	FROM User_votes
	WHERE nomination_id = :id`

	queryDeleteUserVotes = `
	DELETE FROM User_votes
	WHERE id = :id
	RETURNING id, user_id, nomination_id, lkm_id`
)
