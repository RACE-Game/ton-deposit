package deposit

//replenishment

// func (r *Repository) SaveReplenisments(ctx context.Context, replenishments []claim.ClaimReplenishment) (err error) {
// 	// var (
// 	// 	userID int64
// 	// 	amount int64
// 	// )

// 	if len(replenishments) != 1 {
// 		return fmt.Errorf("invalid replenishments count")
// 	}

// 	replenisment := replenishments[0]

// 	query := fmt.Sprintf(`INSERT INTO %s.replenishments
// 	(token, user_id, claim_id, amount,tx_lt,tx_hash,wallet, comment, created_at) VALUES ($1, $2, $3, $4, $5, $6, $7,$8,$9)
// 	returning id`,
// 		r.db.Scheme(),
// 	)

// 	// batch := &pgx.Batch{}
// 	// batch.Queue("insert into ledger(description, amount) values($1, $2)", "q1", 1)
// 	// batch.Queue("insert into ledger(description, amount) values($1, $2)", "q2", 2)
// 	// br := r.db.SendBatch(context.Background(), batch)
// 	// _ = br

// 	_, err = r.db.ExecContext(ctx, query,
// 		replenisment.Token,
// 		replenisment.UserID,
// 		replenisment.ClaimID,
// 		replenisment.Amount,
// 		replenisment.TXLT,
// 		replenisment.TXHash,
// 		replenisment.Wallet,
// 		replenisment.TXComment,
// 		time.Now(),
// 	)
// 	if err != nil {
// 		return fmt.Errorf("can't save claim: %w", err)
// 	}

// 	return nil
// }

// func (r *Repository) GetAll(ctx context.Context) (replanishments []model.ClaimReplenishment, err error) {
// 	q := fmt.Sprintf(`SELECT id, token,claim_id,user_id,wallet,amount, tx_hash,tx_timestamp,created_at FROM %s.replenishments`, r.db.Scheme())

// 	rows, err := r.db.QueryContext(ctx, q)
// 	if err != nil {
// 		return nil, fmt.Errorf("can't get replenishments: %w", err)

// 	}
// 	defer rows.Close()

// 	for rows.Next() {
// 		var replenishment Replenishment
// 		err = rows.Scan(
// 			&replenishment.ID,
// 			&replenishment.Token,
// 			&replenishment.ClaimID,
// 			&replenishment.UserID,
// 			&replenishment.Wallet,
// 			&replenishment.Amount,
// 			&replenishment.TXHash,
// 			&replenishment.TXTimestamp,
// 			&replenishment.CreatedAt,
// 		)
// 		if err != nil {
// 			return nil, fmt.Errorf("can't scan replenishment: %w", err)
// 		}

// 		r := model.ClaimReplenishment{
// 			ID:          replenishment.ID,
// 			Token:       replenishment.Token,
// 			ClaimID:     replenishment.ClaimID,
// 			UserID:      replenishment.UserID,
// 			Wallet:      replenishment.Wallet,
// 			Amount:      replenishment.Amount,
// 			TXHash:      replenishment.TXHash,
// 			TXTimestamp: replenishment.TXTimestamp,
// 			CreatedAt:   replenishment.CreatedAt,
// 		}

// 		replanishments = append(replanishments, r)
// 	}

// 	return
// }

// func (r *Repository) GetLastReplenismentTXLT(ctx context.Context) (lt uint64, err error) {
// 	q := fmt.Sprintf(`SELECT MAX(tx_lt) FROM %s.replenishments`, r.db.Scheme())

// 	t := pgtype.Numeric{}

// 	row := r.db.QueryRowContext(ctx, q)
// 	err = row.Scan(&t)
// 	if err != nil {
// 		return 0, fmt.Errorf("can't get last replenishment tx timestamp: %w", err)
// 	}

// 	if !t.Valid {
// 		return 0, nil
// 	}

// 	if !t.Int.IsUint64() {
// 		return 0, fmt.Errorf("can't convert to uint64")
// 	}

// 	lt = t.Int.Uint64()

// 	return
// }
