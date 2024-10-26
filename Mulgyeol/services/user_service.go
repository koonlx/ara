package service

import (
    "context"
    "encoding/json"
    "ara/mulgyeol/kafka"
    "ara/mulgyeol/model"
)

type UserService interface {
    SendUserData(user *model.User) error
}

type userService struct {
    producer *kafka.KafkaProducer
}

func NewUserService(producer *kafka.KafkaProducer) UserService {
    return &userService{producer: producer}
}

func (s *userService) SendUserData(user *model.User) error {
    data, err := json.Marshal(user)
    if err != nil {
        return err
    }
    return s.producer.SendMessage(context.Background(), []byte(user.Name), data)
}
