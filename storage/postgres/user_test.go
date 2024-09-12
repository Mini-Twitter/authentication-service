package postgres

import (
	_ "github.com/lib/pq"
)

//func TestCreate(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.CreateRequest{
//		Phone:       "1234567890",
//		Email:       "john.doe@example.com",
//		Password:    "password123",
//		FirstName:   "John",
//		LastName:    "Doe",
//		Username:    "johndoe",
//		Nationality: "American",
//		Bio:         "Software Developer",
//	}
//
//	res, err := repo.Create(req)
//	if err != nil {
//		t.Fatalf("Failed to create user: %v", err)
//	}
//
//	if res.Email != req.Email {
//		t.Errorf("Expected email %v, got %v", req.Email, res.Email)
//	}
//
//	fmt.Println("Create response:", res)
//}
//
//// TestGetProfile tests the GetProfile method of UserRepo.
//func TestGetProfile(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	// Generate a valid UUID and ensure it exists in the test database
//	userID := uuid.New().String()
//
//	// Insert test data into the database
//	_, err = db.Exec(`INSERT INTO users (id, phone, email, password) VALUES ($1, $2, $3, $4)`,
//		userID, "1234567890", "test.user@example.com", "password123")
//	if err != nil {
//		t.Fatalf("Failed to insert test user: %v", err)
//	}
//
//	_, err = db.Exec(`INSERT INTO user_profile (user_id, first_name, last_name, username, nationality, bio) VALUES ($1, $2, $3, $4, $5, $6)`,
//		userID, "John", "Doe", "johndoe", "American", "Software Developer")
//	if err != nil {
//		t.Fatalf("Failed to insert user profile: %v", err)
//	}
//
//	// Run the test
//	req := &pb.Id{UserId: userID}
//	res, err := repo.GetProfile(req)
//	if err != nil {
//		t.Fatalf("Failed to get profile: %v", err)
//	}
//
//	if res.UserId != userID {
//		t.Errorf("Expected user ID %v, got %v", userID, res.UserId)
//	}
//
//	fmt.Println("GetProfile response:", res)
//	// Clean up the test data
//	_, err = db.Exec(`DELETE FROM user_profile WHERE user_id = $1`, userID)
//	if err != nil {
//		t.Fatalf("Failed to delete user profile: %v", err)
//	}
//
//	_, err = db.Exec(`DELETE FROM users WHERE id = $1`, userID)
//	if err != nil {
//		t.Fatalf("Failed to delete test user: %v", err)
//	}
//
//}

//// TestUpdateProfile tests the UpdateProfile method of UserRepo.
//func TestUpdateProfile(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.UpdateProfileRequest{
//		UserId:       "some_existing_user_id", // Replace with an actual user ID that exists in your test database
//		FirstName:    "Jane",
//		LastName:     "Smith",
//		Username:     "janesmith",
//		Nationality:  "Canadian",
//		Bio:          "Graphic Designer",
//		ProfileImage: "http://example.com/image.jpg",
//	}
//
//	res, err := repo.UpdateProfile(req)
//	if err != nil {
//		t.Fatalf("Failed to update profile: %v", err)
//	}
//
//	if res.Id != req.UserId {
//		t.Errorf("Expected user ID %v, got %v", req.UserId, res.Id)
//	}
//
//	fmt.Println("UpdateProfile response:", res)
//}
//
//// TestChangePassword tests the ChangePassword method of UserRepo.
//func TestChangePassword(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.ChangePasswordRequest{
//		UserId:          "some_existing_user_id", // Replace with an actual user ID
//		CurrentPassword: "oldpassword",
//		NewPassword:     "newpassword123",
//	}
//
//	res, err := repo.ChangePassword(req)
//	if err != nil {
//		t.Fatalf("Failed to change password: %v", err)
//	}
//
//	if res.Message != "Password updated successfully" {
//		t.Errorf("Expected message %v, got %v", "Password updated successfully", res.Message)
//	}
//
//	fmt.Println("ChangePassword response:", res)
//}
//
//// TestFollow tests the Follow method of UserRepo.
//func TestFollow(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.FollowReq{
//		FollowerId:  "follower_user_id", // Replace with actual user IDs
//		FollowingId: "following_user_id",
//	}
//
//	res, err := repo.Follow(req)
//	if err != nil {
//		t.Fatalf("Failed to follow user: %v", err)
//	}
//
//	if res.FollowerId != req.FollowerId || res.FollowingId != req.FollowingId {
//		t.Errorf("Expected follower ID %v and following ID %v, got follower ID %v and following ID %v",
//			req.FollowerId, req.FollowingId, res.FollowerId, res.FollowingId)
//	}
//
//	fmt.Println("Follow response:", res)
//}
//
//// TestUnfollow tests the Unfollow method of UserRepo.
//func TestUnfollow(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.FollowReq{
//		FollowerId:  "follower_user_id", // Replace with actual user IDs
//		FollowingId: "following_user_id",
//	}
//
//	res, err := repo.Unfollow(req)
//	if err != nil {
//		t.Fatalf("Failed to unfollow user: %v", err)
//	}
//
//	if res.FollowerId != req.FollowerId || res.FollowingId != req.FollowingId {
//		t.Errorf("Expected follower ID %v and following ID %v, got follower ID %v and following ID %v",
//			req.FollowerId, req.FollowingId, res.FollowerId, res.FollowingId)
//	}
//
//	fmt.Println("Unfollow response:", res)
//}
//
//// TestGetUserFollowers tests the GetUserFollowers method of UserRepo.
//func TestGetUserFollowers(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.Id{UserId: "some_existing_user_id"} // Replace with actual user ID
//
//	res, err := repo.GetUserFollowers(req)
//	if err != nil {
//		t.Fatalf("Failed to get user followers: %v", err)
//	}
//
//	fmt.Println("GetUserFollowers response:", res)
//}
//
//// TestGetUserFollows tests the GetUserFollows method of UserRepo.
//func TestGetUserFollows(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.Id{UserId: "some_existing_user_id"} // Replace with actual user ID
//
//	res, err := repo.GetUserFollows(req)
//	if err != nil {
//		t.Fatalf("Failed to get user follows: %v", err)
//	}
//
//	fmt.Println("GetUserFollows response:", res)
//}
//
//// TestMostPopularUser tests the MostPopularUser method of UserRepo.
//func TestMostPopularUser(t *testing.T) {
//	db, err := Connect()
//	if err != nil {
//		t.Fatalf("Failed to connect to database: %v", err)
//	}
//	defer db.Close()
//
//	repo := NewUserRepo(db)
//
//	req := &pb.Void{}
//
//	res, err := repo.MostPopularUser(req)
//	if err != nil {
//		t.Fatalf("Failed to get most popular user: %v", err)
//	}
//
//	fmt.Println("MostPopularUser response:", res)
//}
