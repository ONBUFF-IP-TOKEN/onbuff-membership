package model

import (
	"fmt"

	"github.com/ONBUFF-IP-TOKEN/baseutil/datetime"
	"github.com/ONBUFF-IP-TOKEN/baseutil/log"
	"github.com/ONBUFF-IP-TOKEN/onbuff-membership/rest_server/controllers/context"
)

func (o *DB) GetExistMember(walletAddr string) (*context.Member, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM members WHERE wallet_address='%v'", walletAddr)
	rows, err := o.Mysql.Query(sqlQuery)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	member := context.NewMember()
	for rows.Next() {
		if err := rows.Scan(&member.Id, &member.WalletAddr, &member.Email, &member.WalletType, &member.CreateTs, &member.NickName,
			&member.ProfileImg, &member.ActivateState); err != nil {
			log.Error(err)
		}
	}

	return member, err
}

func (o *DB) InsertMember(memberInfo *context.RegisterMember) (int64, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO members(wallet_address, email, wallet_type, create_ts, nickname, profile_img, activate_state) " +
		"VALUES (?,?,?,?,?,?,?)")

	result, err := o.Mysql.PrepareAndExec(sqlQuery, memberInfo.WalletAuth.WalletAddr, memberInfo.Email, memberInfo.WalletType,
		datetime.GetTS2MilliSec(), memberInfo.NickName, memberInfo.ProfileImg, memberInfo.ActivateState)
	if err != nil {
		log.Error(err)
		return -1, err
	}

	insertId, err := result.LastInsertId()
	if err != nil {
		log.Error(err)
		return -1, err
	}
	log.Debug("InsertMember id:", insertId)
	return insertId, nil
}

func (o *DB) UpdateMember(memberInfo *context.Member) (int64, error) {
	sqlQuery := "UPDATE members set wallet_address=?,email=?,wallet_type=?,nickname=?,profile_img=?,activate_state=? WHERE id=?"

	result, err := o.Mysql.PrepareAndExec(sqlQuery, memberInfo.WalletAddr, memberInfo.Email, memberInfo.WalletType, memberInfo.NickName,
		memberInfo.ProfileImg, memberInfo.ActivateState, memberInfo.Id)
	if err != nil {
		log.Error(err)
		return 0, err
	}
	cnt, err := result.RowsAffected()
	if err != nil {
		log.Error(err)
		return 0, err
	}

	return cnt, nil
}
