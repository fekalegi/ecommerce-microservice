package user_test

import (
	"ecommerce-microservice/user/common/dto"
	"ecommerce-microservice/user/domain"
	mock_repository "ecommerce-microservice/user/mocks/repository"
	"ecommerce-microservice/user/usecase"
	"ecommerce-microservice/user/usecase/user"
	"errors"
	"github.com/fekalegi/custom-package/authentications/middlewares/common/interfaces"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errSomething = errors.New("something error")

var _ = Describe("UserService", func() {
	var (
		mockCtrl  *gomock.Controller
		userUC    usecase.UserService
		repo      *mock_repository.MockUserRepository
		mockUser  domain.User
		mockUsers []domain.User
		helper    interfaces.HelperInterface
		mockUID   uuid.UUID
		//mockUID2  uuid.UUID
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockUserRepository(mockCtrl)
		userUC = user.NewUserImplementation(repo, helper)
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a3")
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a4")

		mockUser = domain.User{
			ID:       1,
			Username: "feka",
			Password: "feka",
			FullName: "Feka",
			AuthUUID: mockUID,
			RoleID:   1,
			Role: domain.Role{
				ID:    1,
				Name:  "superadmin",
				Level: 1,
			},
		}

		mockUsers = []domain.User{
			{
				ID:       1,
				Username: "feka",
				Password: "feka",
				FullName: "Feka",
				AuthUUID: mockUID,
				RoleID:   1,
				Role: domain.Role{
					ID:    1,
					Name:  "superadmin",
					Level: 1,
				},
			}, {
				ID:       2,
				Username: "feka",
				Password: "feka",
				FullName: "Feka",
				AuthUUID: mockUID,
				RoleID:   2,
				Role: domain.Role{
					ID:    2,
					Name:  "admin",
					Level: 2,
				},
			},
		}

	})

	Describe("AddUser", func() {
		mockRequest := dto.RequestUser{
			Username: "feka",
			Password: "feka",
			FullName: "Feka",
			RoleID:   1,
		}
		It("return success", func() {
			repo.EXPECT().AddUser(gomock.Any()).Return(&mockUser, nil)
			resp, err := userUC.AddUser(mockRequest)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(201))
		})

		It("return error", func() {
			repo.EXPECT().AddUser(gomock.Any()).Return(nil, errSomething)
			_, err := userUC.AddUser(mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("CheckLogin", func() {
		It("return success", func() {
			repo.EXPECT().CheckLogin(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			repo.EXPECT().UpdateAuthUUID(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := userUC.CheckLogin("feka", "feka")
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return error", func() {
			repo.EXPECT().CheckLogin(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			repo.EXPECT().UpdateAuthUUID(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := userUC.CheckLogin("feka", "feka")
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteAuth", func() {
		It("return success", func() {
			repo.EXPECT().FindUserByIDAndAuthID(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			repo.EXPECT().UpdateAuthUUID(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := userUC.DeleteAuth(1, mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return success", func() {
			repo.EXPECT().FindUserByIDAndAuthID(gomock.Any(), gomock.Any()).Return(&mockUser, nil)
			repo.EXPECT().UpdateAuthUUID(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := userUC.DeleteAuth(1, mockUID)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteUserByID", func() {
		It("return success", func() {
			repo.EXPECT().FindUserByID(gomock.Any()).Return(&mockUser, nil)
			repo.EXPECT().DeleteUser(gomock.Any()).Return(nil)
			resp, err := userUC.DeleteUser(1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return error", func() {
			repo.EXPECT().FindUserByID(gomock.Any()).Return(&mockUser, nil)
			repo.EXPECT().DeleteUser(gomock.Any()).Return(errSomething)
			_, err := userUC.DeleteUser(1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchAllUser", func() {
		It("return success", func() {
			repo.EXPECT().FindAllUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockUsers, int64(2), nil)
			resp, err := userUC.FetchAllUser(dto.RequestPagination{})
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(mockUsers))
		})

		It("return error", func() {
			repo.EXPECT().FindAllUser(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, int64(0), errSomething)
			_, err := userUC.FetchAllUser(dto.RequestPagination{})
			Expect(err).Should(HaveOccurred())
		})
	})
})
