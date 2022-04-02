package wishlist_test

import (
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	mock_repository "ecommerce-microservice/product/mocks/repository"
	"ecommerce-microservice/product/usecase"
	"ecommerce-microservice/product/usecase/wishlist"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errSomething = errors.New("something error")

var _ = Describe("WishlistWishlist", func() {
	var (
		mockCtrl      *gomock.Controller
		wishlistUC    usecase.WishlistService
		repo          *mock_repository.MockWishlistRepository
		mockWishlist  domain.Wishlist
		mockWishlists []domain.Wishlist
		mockUID       uuid.UUID
		mockUID2      uuid.UUID
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockWishlistRepository(mockCtrl)
		wishlistUC = wishlist.NewWishlistRepository(repo)
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a3")
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a4")

		mockWishlist = domain.Wishlist{
			ID:        mockUID,
			ProductID: mockUID,
			UserID:    1,
		}

		mockWishlists = []domain.Wishlist{
			{
				ID:        mockUID,
				ProductID: mockUID,
				UserID:    1,
			}, {
				ID:        mockUID2,
				ProductID: mockUID2,
				UserID:    1,
			},
		}

	})

	Describe("CreateWishlistWishlistCommand", func() {
		It("return succeed", func() {
			request := dto.RequestWishlist{
				ID:        mockUID,
				ProductID: mockUID,
			}
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateWishlist(gomock.Any()).Return(nil)
			_, err := wishlistUC.CreateWishlistCommand(request, 1, 1)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			request := dto.RequestWishlist{
				ID:        mockUID,
				ProductID: mockUID,
			}
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateWishlist(gomock.Any()).Return(errSomething)
			_, err := wishlistUC.CreateWishlistCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchAllWishlist", func() {
		mockRequest := dto.RequestPagination{
			Filter:     "",
			PageSize:   0,
			PageNumber: 0,
		}
		mockCount := int64(1)

		It("return succeed", func() {
			repo.EXPECT().FetchAllWishlistByUserID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockWishlists, mockCount, nil)
			resp, err := wishlistUC.FetchAllWishlistByUserID(mockRequest, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(mockWishlists))
		})

		It("return error", func() {
			repo.EXPECT().FetchAllWishlistByUserID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockWishlists, mockCount, errSomething)
			_, err := wishlistUC.FetchAllWishlistByUserID(mockRequest, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchWishlistByID", func() {
		It("return succeed", func() {
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(&mockWishlist, nil)
			resp, err := wishlistUC.FetchWishlistQuery(mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(&mockWishlist))
		})

		It("return not found", func() {
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(nil, nil)
			resp, err := wishlistUC.FetchWishlistQuery(mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(nil, errSomething)
			_, err := wishlistUC.FetchWishlistQuery(mockUID)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteWishlistWishlist", func() {
		It("return succeed", func() {
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(&mockWishlist, nil)
			repo.EXPECT().HardDeleteWishlist(gomock.Any()).Return(nil)
			resp, err := wishlistUC.DeleteWishlistCommand(mockUID, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return not found", func() {
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().HardDeleteWishlist(gomock.Any()).Return(nil)
			resp, err := wishlistUC.DeleteWishlistCommand(mockUID, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchWishlistByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().HardDeleteWishlist(gomock.Any()).Return(errSomething)
			_, err := wishlistUC.DeleteWishlistCommand(mockUID, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})
})
