
// Package userAccount ユーザーアカウント
package userAccount

import (
	
"google.golang.org/protobuf/types/known/timestamppb"
"google.golang.org/protobuf/types/known/timestamppb"
"google.golang.org/protobuf/types/known/timestamppb"
"google.golang.org/protobuf/types/known/timestamppb"
)

type UserAccounts []*UserAccount

func NewUserAccount() *UserAccount {
			return &UserAccount{}
		}

		func NewUserAccounts() UserAccounts {
			return UserAccounts{}
		}

		func SetUserAccount(userId string,name string,password string,loginAt *timestamppb.Timestamp,logoutAt *timestamppb.Timestamp) *UserAccount {
			return &UserAccount{
				UserId: userId,
Name: name,
Password: password,
LoginAt: loginAt,
LogoutAt: logoutAt,
			}
		}
		
