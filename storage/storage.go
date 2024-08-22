package storage

import (
	pb "auth-service/genproto/user"
	"auth-service/pkg/models"
)

type AuthStorage interface {
	Register(in models.RegisterRequest) (models.RegisterResponse, error)
	LoginEmail(in models.LoginEmailRequest) (models.LoginEmailResponse, error)
	LoginUsername(in models.LoginUsernameRequest) (models.LoginUsernameResponse, error)
}

type UserStorage interface {
	Create(in *pb.CreateRequest) (*pb.UserResponse, error)
	GetProfile(in *pb.Id) (*pb.GetProfileResponse, error)
	UpdateProfile(in *pb.UpdateProfileRequest) (*pb.UserResponse, error)
	ChangePassword(in *pb.ChangePasswordRequest) (*pb.ChangePasswordResponse, error)
	ChangeProfileImage(in *pb.URL) (*pb.Void, error)
	FetchUsers(in *pb.Filter) (*pb.UserResponses, error)
	Follow(in *pb.Ids) (*pb.Void, error)
	UnFollow(in *pb.Ids) (*pb.Void, error)
	PostAdd(in *pb.Id) (*pb.Void, error)
	ListOfFollowing(in *pb.Id) (*pb.Followings, error)
	ListOfFollowers(in *pb.Id) (*pb.Followers, error)
	PostDelete(in *pb.Id) (*pb.Void, error)
	DeleteUser(in *pb.Id) (*pb.Void, error)
}
