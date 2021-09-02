package model

import (
	"database/sql"
	"encoding/json"
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

	var tos sql.NullString

	member := context.NewMember()
	for rows.Next() {
		if err := rows.Scan(&member.Id, &member.WalletAddr, &member.Email, &member.WalletType, &member.CreateTs, &member.NickName,
			&member.ProfileImg, &member.ActivateState, &tos); err != nil {
			log.Error(err)
		} else {
			getTos := context.TermsOfService{}
			json.Unmarshal([]byte(tos.String), &getTos)
			member.TermsOfService = getTos
		}
	}

	return member, err
}

func (o *DB) InsertMember(memberInfo *context.RegisterMember) (int64, error) {
	sqlQuery := fmt.Sprintf("INSERT INTO members(wallet_address, email, wallet_type, create_ts, nickname, profile_img, activate_state,terms_of_service ) " +
		"VALUES (?,?,?,?,?,?,?,?)")

	tos, _ := json.Marshal(memberInfo.TermsOfService)

	result, err := o.Mysql.PrepareAndExec(sqlQuery, memberInfo.WalletAuth.WalletAddr, memberInfo.Email, memberInfo.WalletType,
		datetime.GetTS2MilliSec(), memberInfo.NickName, memberInfo.ProfileImg, memberInfo.ActivateState, tos)
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
	// When updating, the terms and conditions are not processed.
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

func (o *DB) GetExistMemberByNickEmail(nickname, email string) (*context.Member, error) {
	sqlQuery := "SELECT * FROM members WHERE "
	if len(email) > 0 && len(nickname) == 0 {
		sqlQuery += "email='" + email + "'"
	} else if len(email) == 0 && len(nickname) > 0 {
		sqlQuery += "nickname='" + nickname + "'"
	} else if len(email) > 0 && len(nickname) > 0 {
		sqlQuery += "email='" + email + "' OR nickname='" + nickname + "'"
	} else {
		return context.NewMember(), nil
	}

	rows, err := o.Mysql.Query(sqlQuery)

	if err != nil {
		log.Error(err)
		return nil, err
	}

	defer rows.Close()

	var tos sql.NullString
	member := context.NewMember()
	for rows.Next() {
		if err := rows.Scan(&member.Id, &member.WalletAddr, &member.Email, &member.WalletType, &member.CreateTs, &member.NickName,
			&member.ProfileImg, &member.ActivateState, &tos); err != nil {
			log.Error(err)
		} else {
			getTos := context.TermsOfService{}
			json.Unmarshal([]byte(tos.String), &getTos)
			member.TermsOfService = getTos
		}
	}

	return member, err
}

func (o *DB) UpdateMemberWithdraw(memberInfo *context.Member) (int64, error) {
	// When updating, the terms and conditions are not processed.
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

func (o *DB) DeleteMember(memberInfo *context.Member) error {
	sqlQuery := "DELETE FROM members WHERE wallet_address=?"
	result, err := o.Mysql.PrepareAndExec(sqlQuery, memberInfo.WalletAddr)
	if err != nil {
		log.Error(err)
		return err
	}
	if _, err = result.RowsAffected(); err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (o *DB) GetMemberList(memberList *context.MemberList) (*[]context.Member, int64, error) {
	sqlQuery := fmt.Sprintf("SELECT * FROM members ORDER BY id DESC LIMIT %v,%v", memberList.PageSize*memberList.PageOffset, memberList.PageSize)

	rows, err := o.Mysql.Query(sqlQuery)
	if err != nil {
		log.Error(err)
		return nil, 0, err
	}

	defer rows.Close()
	members := []context.Member{}
	var tos sql.NullString

	for rows.Next() {
		member := context.Member{}
		if err := rows.Scan(&member.Id, &member.WalletAddr, &member.Email, &member.WalletType, &member.CreateTs, &member.NickName,
			&member.ProfileImg, &member.ActivateState, &tos); err != nil {
			log.Error(err)
		} else {
			getTos := context.TermsOfService{}
			json.Unmarshal([]byte(tos.String), &getTos)
			member.TermsOfService = getTos

			members = append(members, member)
		}
	}

	total, _ := o.GetTotalMember()
	return &members, total, nil
}

func (o *DB) GetTotalMember() (int64, error) {
	var count int64
	query := fmt.Sprintf("SELECT COUNT(*) as count FROM members")
	err := o.Mysql.QueryRow(query, &count)

	if err != nil {
		log.Error(err)
		return count, err
	}

	return count, nil
}
