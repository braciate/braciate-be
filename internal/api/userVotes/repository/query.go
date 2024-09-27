package votesRepository

const (
	queryCreateUserVotes = `
	INSERT INTO User_votes (id, user_id, lkm_id, nomination_id)
	VALUES (:id, :user_id, :lkm_id, :nomination_id)
	RETURNING id, user_id, lkm_id, nomination_id
	`

	queryGetUserVoteFromNominationID = `
	SELECT id, user_id, nomination_id, lkm_id
	FROM UserVotes
	WHERE nomination_id = :id`

	queryUpdateUserVote = `
	UPDATE UserVotes
	SET nomination_id = :nomination_id, lkm_id =:lkm_id
	RETURNING id, user_id, nomination_id, lkm_id`

	queryDeleteUserVotes = `
	DELETE FROM UserVotes
	WHERE id = :id
	RETURNING id, user_id, nomination_id, lkm_id`
)
