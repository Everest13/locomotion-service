package user_server

import (
	"google.golang.org/protobuf/types/known/timestamppb"
	"locomotion-service/internal/service/user"
	desc "locomotion-service/pkg/user-service"
	"time"
)

func convertUserToPB(user *user.User) *desc.User {
	return &desc.User{
		Id:         user.ID,
		FirstName:  user.FirstName,
		SecondName: user.SecondName,
		Patronymic: *user.Patronymic,
		Age:        user.Age,
		Sex:        convertSexToPb(user.Sex),
		CreatedAt:  convertTimeToTimestamp(user.CreatedAt),
		DateBirth:  convertTimeToTimestamp(user.DateBirth),
	}
}

func convertUserFromPB(userPb *desc.UserCreate) *user.User {
	patronymic := userPb.GetPatronymic()
	return &user.User{
		FirstName:  userPb.GetFirstName(),
		SecondName: userPb.GetSecondName(),
		Patronymic: &patronymic,
		Age:        userPb.GetAge(),
		Sex:        convertSexFromPb(userPb.GetSex()),
		DateBirth:  *convertTimestampToTimePointer(userPb.GetDateBirth()),
	}
}

func convertSexToPb(sex user.SexType) desc.UserSex {
	switch sex {
	case user.FemaleSexType:
		return desc.UserSex_FEMALE
	case user.MaleSexType:
		return desc.UserSex_MALE
	default:
		return desc.UserSex_UNKNOWN
	}
}

func convertSexFromPb(sex desc.UserSex) user.SexType {
	switch sex {
	case desc.UserSex_FEMALE:
		return user.FemaleSexType
	case desc.UserSex_MALE:
		return user.MaleSexType
	default:
		return user.UnknownSexType
	}
}

func convertTimeToTimestamp(input time.Time) *timestamppb.Timestamp {
	return &timestamppb.Timestamp{
		Seconds: input.Unix(),
		Nanos:   int32(input.Nanosecond()),
	}
}

func convertTimestampToTimePointer(input *timestamppb.Timestamp) *time.Time {
	if input == nil {
		return nil
	}

	result := input.AsTime()
	return &result
}
