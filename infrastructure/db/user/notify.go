package user

import (
	"context"
	"fmt"
)

func (r *Repository) GetAllUserID(ctx context.Context) ([]int64, error) {
	q := `SELECT DISTINCT user_id FROM ( 
SELECT  user_id FROM memepolis.gameovers 
UNION ALL    
SELECT  user_id 
FROM memepolis.lifes    
UNION ALL 
SELECT DISTINCT user_id FROM memepolis.scores 
UNION ALL 
SELECT DISTINCT referrer_user_id as user_id FROM memepolis.referals 
UNION ALL       
SELECT  referal_user_id  as user_id FROM memepolis.referals
UNION ALL    SELECT  telegram_id 
FROM memepolis.users
) AS combined_results 
order by user_id;`
	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to get all user id: %w", err)
	}
	defer rows.Close()

	var userIDs []int64

	for rows.Next() {
		var userID int64
		err = rows.Scan(&userID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user id: %w", err)
		}
		userIDs = append(userIDs, userID)
	}

	return userIDs, nil
}

func (r *Repository) SaveNotifyResult(ctx context.Context, userID int64, appErr string) error {
	q := fmt.Sprintf(`INSERT INTO %s.notification_results (user_id, result) 
	VALUES ($1, $2) ON CONFLICT (user_id) DO UPDATE SET result = $2`,
		r.db.Scheme())

	_, err := r.db.ExecContext(ctx, q, userID, appErr)
	if err != nil {
		return fmt.Errorf("failed to save notification result: %w", err)
	}

	return nil

}

func (r *Repository) GetFailedNotify(ctx context.Context) (map[int64]struct{}, error) {
	q := fmt.Sprintf(`SELECT user_id FROM %s.notification_results WHERE result IS NOT NULL`,
		r.db.Scheme())

	rows, err := r.db.QueryContext(ctx, q)
	if err != nil {
		return nil, fmt.Errorf("failed to get failed notify: %w", err)
	}
	defer rows.Close()

	failedMap := make(map[int64]struct{})

	for rows.Next() {
		var userID int64
		err = rows.Scan(&userID)
		if err != nil {
			return nil, fmt.Errorf("failed to scan failed notify: %w", err)
		}
		failedMap[userID] = struct{}{}
	}

	return failedMap, nil
}
