package product_test

import (
	"ecommerce-microservice/product/common/dto"
	"ecommerce-microservice/product/domain"
	mock_repository "ecommerce-microservice/product/mocks/repository"
	"ecommerce-microservice/product/usecase"
	"ecommerce-microservice/product/usecase/product"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errSomething = errors.New("something error")

var _ = Describe("ProductCategory", func() {
	var (
		mockCtrl     *gomock.Controller
		productUC    usecase.ProductService
		repo         *mock_repository.MockProductRepository
		mockProduct  domain.Product
		mockProducts []domain.Product
		mockUID      uuid.UUID
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockProductRepository(mockCtrl)
		productUC = product.NewProductService(repo)
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a3")

		mockProduct = domain.Product{
			ID:          mockUID,
			Name:        "Baju",
			Description: "Ini Baju Bagus",
			Price:       100000,
			Stock:       10,
			CategoryID:  uuid.UUID{},
			UserID:      1,
		}

		mockProducts = []domain.Product{
			{
				ID:          mockUID,
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
				UserID:      1,
			}, {
				ID:          mockUID,
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
				UserID:      1,
			},
		}

	})

	Describe("CreateProductCategoryCommand", func() {
		It("return succeed", func() {
			request := dto.RequestProduct{
				ID:          mockUID,
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
			}
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateProduct(gomock.Any()).Return(nil)
			_, err := productUC.CreateProductCommand(request, 1, 1)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			request := dto.RequestProduct{
				ID:          mockUID,
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
			}
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateProduct(gomock.Any()).Return(errSomething)
			_, err := productUC.CreateProductCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchAllProduct", func() {
		mockRequest := dto.RequestPagination{
			Filter:     "",
			PageSize:   0,
			PageNumber: 0,
		}
		mockCount := int64(1)

		It("return succeed", func() {
			repo.EXPECT().FetchAllProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(mockProducts, mockCount, nil)
			resp, err := productUC.FetchAllProductQuery(mockRequest)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(mockProducts))
		})

		It("return error", func() {
			repo.EXPECT().FetchAllProduct(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil, mockCount, errSomething)
			_, err := productUC.FetchAllProductQuery(mockRequest)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchProductByID", func() {
		It("return succeed", func() {
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(&mockProduct, nil)
			resp, err := productUC.FetchProductQuery(mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(&mockProduct))
		})

		It("return not found", func() {
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, nil)
			resp, err := productUC.FetchProductQuery(mockUID)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, errSomething)
			_, err := productUC.FetchProductQuery(mockUID)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("UpdateProductCategoryCommand", func() {
		It("return succeed", func() {
			request := dto.RequestUpdateProduct{
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
			}
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(&mockProduct, nil)
			repo.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := productUC.UpdateProductCommand(mockUID, request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return not found", func() {
			request := dto.RequestUpdateProduct{
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
			}
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := productUC.UpdateProductCommand(mockUID, request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			request := dto.RequestUpdateProduct{
				Name:        "Baju",
				Description: "Ini Baju Bagus",
				Price:       100000,
				Stock:       10,
				CategoryID:  uuid.UUID{},
			}
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateProduct(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := productUC.UpdateProductCommand(mockUID, request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteProductCategory", func() {
		It("return succeed", func() {
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(&mockProduct, nil)
			repo.EXPECT().DeleteProduct(gomock.Any()).Return(nil)
			resp, err := productUC.DeleteProductCommand(mockUID, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return not found", func() {
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().DeleteProduct(gomock.Any()).Return(nil)
			resp, err := productUC.DeleteProductCommand(mockUID, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchProductByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().DeleteProduct(gomock.Any()).Return(errSomething)
			_, err := productUC.DeleteProductCommand(mockUID, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})
})
