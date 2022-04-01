package category_test

import (
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	mock_repository "ecommerce-microservice/product/mocks/repository"
	"ecommerce-microservice/product/usecase"
	"ecommerce-microservice/product/usecase/category"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errSomething = errors.New("something error")

var _ = Describe("CategoryCategory", func() {
	var (
		mockCtrl      *gomock.Controller
		categoryUC    usecase.CategoryService
		repo          *mock_repository.MockCategoryRepository
		mockCategory  domain.Category
		mockCategorys []domain.Category
		mockUID       uuid.UUID
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockCategoryRepository(mockCtrl)
		categoryUC = category.NewCategoryService(repo)
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a3")

		mockCategory = domain.Category{
			ID:   mockUID,
			Name: "Baju",
		}

		mockCategorys = []domain.Category{
			{
				ID:   mockUID,
				Name: "Baju",
			}, {
				ID:   mockUID,
				Name: "Baju",
			},
		}

	})

	Describe("CreateCategoryCategoryCommand", func() {
		It("return succeed", func() {
			request := dto.RequestCategory{
				ID:   mockUID,
				Name: "Baju",
			}
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateCategory(gomock.Any()).Return(nil)
			_, err := categoryUC.CreateCategoryCommand(request, 1, 1)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			request := dto.RequestCategory{
				ID:   mockUID,
				Name: "Baju",
			}
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateCategory(gomock.Any()).Return(errSomething)
			_, err := categoryUC.CreateCategoryCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchAllCategory", func() {
		mockRequest := dto.RequestPagination{
			Filter:     "",
			PageSize:   0,
			PageNumber: 0,
		}
		mockCount := int64(1)

		It("return succeed", func() {
			repo.EXPECT().FetchAllCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockCategorys, mockCount, nil)
			resp, err := categoryUC.FetchAllCategoryQuery(mockRequest)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(mockCategorys))
		})

		It("return error", func() {
			repo.EXPECT().FetchAllCategory(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, mockCount, errSomething)
			_, err := categoryUC.FetchAllCategoryQuery(mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchCategoryByID", func() {
		It("return succeed", func() {
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(&mockCategory, nil)
			resp, err := categoryUC.FetchCategoryQuery(mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(&mockCategory))
		})

		It("return not found", func() {
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, nil)
			resp, err := categoryUC.FetchCategoryQuery(mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, errSomething)
			_, err := categoryUC.FetchCategoryQuery(mockUID)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("UpdateCategoryCategoryCommand", func() {
		It("return succeed", func() {
			request := dto.RequestUpdateCategory{
				Name: "Baju",
			}
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(&mockCategory, nil)
			repo.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := categoryUC.UpdateCategoryCommand(mockUID, request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return not found", func() {
			request := dto.RequestUpdateCategory{
				Name: "Baju",
			}
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := categoryUC.UpdateCategoryCommand(mockUID, request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			request := dto.RequestUpdateCategory{
				Name: "Baju",
			}
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateCategory(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := categoryUC.UpdateCategoryCommand(mockUID, request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteCategoryCategory", func() {
		It("return succeed", func() {
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(&mockCategory, nil)
			repo.EXPECT().DeleteCategory(gomock.Any()).Return(nil)
			resp, err := categoryUC.DeleteCategoryCommand(mockUID, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return not found", func() {
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().DeleteCategory(gomock.Any()).Return(nil)
			resp, err := categoryUC.DeleteCategoryCommand(mockUID, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchCategoryByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().DeleteCategory(gomock.Any()).Return(errSomething)
			_, err := categoryUC.DeleteCategoryCommand(mockUID, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})
})
