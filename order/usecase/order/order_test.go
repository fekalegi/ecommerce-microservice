package order_test

import (
	"ecommerce-microservice/order/common/dto"
	"ecommerce-microservice/order/domain"
	mock_repository "ecommerce-microservice/order/mocks/repository"
	mock_service "ecommerce-microservice/order/mocks/services"
	"ecommerce-microservice/order/usecase"
	"ecommerce-microservice/order/usecase/order"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var errSomething = errors.New("something error")

var _ = Describe("OrderCategory", func() {
	var (
		mockCtrl   *gomock.Controller
		orderUC    usecase.OrderService
		repo       *mock_repository.MockOrderRepository
		producer   *mock_service.MockProducerService
		mockOrder  domain.Order
		mockOrders []domain.Order
		mockUID    uuid.UUID
		mockRating domain.SellerRating
	)

	BeforeEach(func() {
		mockCtrl = gomock.NewController(GinkgoT())
		mockCtrl.Finish()
		repo = mock_repository.NewMockOrderRepository(mockCtrl)
		orderUC = order.NewOrderService(repo, producer)
		mockUID, _ = uuid.Parse("73932cef-90ee-4299-9fa0-5b5c0326e4a3")

		mockOrder = domain.Order{
			ID:                 1,
			UserID:             1,
			SellerID:           1,
			TotalOrder:         100000,
			Status:             1,
			IsCanceled:         false,
			CancellationStatus: 0,
			Products: []domain.OrderProduct{
				{
					ID:         1,
					OrderID:    1,
					ProductID:  mockUID,
					Quantity:   2,
					Price:      50000,
					TotalPrice: 100000,
				},
			},
		}

		mockOrders = []domain.Order{
			{
				ID:                 1,
				UserID:             1,
				SellerID:           1,
				TotalOrder:         100000,
				Status:             1,
				IsCanceled:         false,
				CancellationStatus: 0,
				Products: []domain.OrderProduct{
					{
						ID:         1,
						OrderID:    1,
						ProductID:  mockUID,
						Quantity:   2,
						Price:      50000,
						TotalPrice: 100000,
					},
				},
			}, {
				ID:                 2,
				UserID:             2,
				SellerID:           2,
				TotalOrder:         150000,
				Status:             1,
				IsCanceled:         false,
				CancellationStatus: 0,
				Products: []domain.OrderProduct{
					{
						ID:         2,
						OrderID:    2,
						ProductID:  mockUID,
						Quantity:   3,
						Price:      50000,
						TotalPrice: 100000,
					},
				},
			},
		}

		mockRating = domain.SellerRating{
			SellerID: 1,
			One:      1,
			Two:      2,
			Three:    3,
			Four:     4,
			Five:     5,
		}

	})

	Describe("CreateInitialOrderCommand", func() {
		It("return succeed", func() {
			request := dto.RequestInitialOrder{
				SellerID: 1,
				Products: dto.RequestOrderProduct{
					ProductID: mockUID,
					Quantity:  2,
					Price:     500000,
				},
			}
			repo.EXPECT().CreateInitialOrder(gomock.Any(), gomock.Any()).Return(nil)
			_, err := orderUC.CreateInitialOrderCommand(request, 1, 1)
			Expect(err).Should(Succeed())
		})

		It("return error", func() {
			request := dto.RequestInitialOrder{
				SellerID: 1,
				Products: dto.RequestOrderProduct{
					ProductID: mockUID,
					Quantity:  2,
					Price:     500000,
				},
			}
			repo.EXPECT().CreateInitialOrder(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := orderUC.CreateInitialOrderCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("AddOrUpdateProductsCommand", func() {
		It("return succeed on create", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return succeed on update", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().UpdateOrderProduct(gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return order not found", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error on FetchOrderByID", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateOrderProduct(gomock.Any()).Return(nil)
			_, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on FindOrderProductByOrderIDAndProductID", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateOrderProduct(gomock.Any(), gomock.Any()).Return(nil)
			_, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on create", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(nil, nil)
			repo.EXPECT().CreateOrderProduct(gomock.Any()).Return(errSomething)
			_, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on update", func() {
			request := dto.RequestOrderProduct{
				ProductID: mockUID,
				Quantity:  2,
				Price:     50000,
			}
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindOrderProductByOrderIDAndProductID(gomock.Any(), gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().UpdateOrderProduct(gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := orderUC.AddOrUpdateOrderProductsCommand(request, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteOrderCommand", func() {
		It("return success", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrder(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderCommand(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return unauthorized user is not the one ordered", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrder(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderCommand(1, 5, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrder(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderCommand(1, 1, 6)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return order not found", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().DeleteOrder(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderCommand(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error on FetchOrderByID", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().DeleteOrder(gomock.Any()).Return(nil)
			_, err := orderUC.DeleteOrderCommand(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on DeleteOrder", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrder(gomock.Any()).Return(errSomething)
			_, err := orderUC.DeleteOrderCommand(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("DeleteOrderProductCommand", func() {
		It("return success", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderProductCommand(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return not found on FetchOrderProductID", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderProductCommand(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return not found on FetchOrderProductID", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderProductCommand(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return unauthorized on user not the one order", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderProductCommand(1, 5, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			resp, err := orderUC.DeleteOrderProductCommand(1, 1, 7)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return error on FetchOrderProductID", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			_, err := orderUC.DeleteOrderProductCommand(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on FetchOrderByID", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, errSomething)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(nil)
			_, err := orderUC.DeleteOrderProductCommand(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on DeleteOrderProduct", func() {
			repo.EXPECT().FetchOrderProductID(gomock.Any()).Return(&mockOrder.Products[0], nil)
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().DeleteOrderProduct(gomock.Any()).Return(errSomething)
			_, err := orderUC.DeleteOrderProductCommand(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchAllProduct", func() {
		mockRequest := dto.RequestFetchOrderByStatus{
			Status:     1,
			PageSize:   0,
			PageNumber: 0,
		}
		mockCount := int64(1)

		It("return succeed", func() {
			repo.EXPECT().FetchAllOrderByUserID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockOrders, mockCount, nil)
			resp, err := orderUC.FetchAllOrderQuery(mockRequest, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(mockOrders))
		})

		It("return error", func() {
			repo.EXPECT().FetchAllOrderByUserID(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).Return(mockOrders, mockCount, errSomething)
			_, err := orderUC.FetchAllOrderQuery(mockRequest, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("FetchOrderQuery", func() {

		It("return succeed", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			resp, err := orderUC.FetchOrderQuery(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Data).Should(Equal(&mockOrder))
		})

		It("return unauthorized, user not the one order", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			resp, err := orderUC.FetchOrderQuery(1, 6, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized, user not the one order", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			resp, err := orderUC.FetchOrderQuery(1, 6, 8)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return not found", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, nil)
			resp, err := orderUC.FetchOrderQuery(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(404))
		})

		It("return error", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			_, err := orderUC.FetchOrderQuery(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("UpdateStatusOrderToWaitForPayment", func() {
		It("return success", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToWaitForPayment(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return unauthorized User not the one order", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToWaitForPayment(1, 3, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToWaitForPayment(1, 3, 5)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return error on FetchOrderByID", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			_, err := orderUC.UpdateStatusOrderToWaitForPayment(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on UpdateOrderStatus", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := orderUC.UpdateStatusOrderToWaitForPayment(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("UpdateStatusOrderToItemsBeingShipped", func() {
		It("return success", func() {
			mockOrder.Status = 3
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToItemsBeingShipped(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return unauthorized User not the one order", func() {
			mockOrder.Status = 3
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToItemsBeingShipped(1, 4, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized", func() {
			mockOrder.Status = 3
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToItemsBeingShipped(1, 4, 5)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return error on FetchOrderByID", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			_, err := orderUC.UpdateStatusOrderToItemsBeingShipped(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on UpdateOrderStatus", func() {
			mockOrder.Status = 3
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := orderUC.UpdateStatusOrderToItemsBeingShipped(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("UpdateStatusOrderToItemsHaveArrived", func() {
		It("return success", func() {
			mockOrder.Status = 4
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToItemsHaveArrived(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return unauthorized User not the one order", func() {
			mockOrder.Status = 4
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToItemsHaveArrived(1, 3, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized", func() {
			mockOrder.Status = 4
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToItemsHaveArrived(1, 3, 5)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return error on FetchOrderByID", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			_, err := orderUC.UpdateStatusOrderToItemsHaveArrived(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on UpdateOrderStatus", func() {
			mockOrder.Status = 4
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := orderUC.UpdateStatusOrderToItemsHaveArrived(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("UpdateStatusOrderToFinished", func() {
		It("return success", func() {
			mockOrder.Status = 5
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToFinished(1, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return unauthorized User not the one order", func() {
			mockOrder.Status = 5
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToFinished(1, 3, 4)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return unauthorized", func() {
			mockOrder.Status = 5
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			resp, err := orderUC.UpdateStatusOrderToFinished(1, 3, 5)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(401))
		})

		It("return error on FetchOrderByID", func() {
			mockOrder.Status = 5
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
			_, err := orderUC.UpdateStatusOrderToFinished(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on UpdateOrderStatus", func() {
			mockOrder.Status = 5
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().UpdateOrderStatus(gomock.Any(), gomock.Any(), gomock.Any()).Return(errSomething)
			_, err := orderUC.UpdateStatusOrderToFinished(1, 1, 1)
			Expect(err).Should(HaveOccurred())
		})
	})

	Describe("AddRatingSeller", func() {
		mockRequest := dto.RequestRatingSeller{Rating: 4}
		It("return success on CreateRatingSeller", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindRatingSellerBySellerID(gomock.Any()).Return(domain.SellerRating{}, nil)
			repo.EXPECT().CreateRatingSeller(gomock.Any()).Return(nil)
			resp, err := orderUC.AddRatingSeller(mockRequest, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return success on UpdateRatingSeller", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindRatingSellerBySellerID(gomock.Any()).Return(mockRating, nil)
			repo.EXPECT().UpdateRatingSeller(gomock.Any()).Return(nil)
			resp, err := orderUC.AddRatingSeller(mockRequest, 1, 1)
			Expect(err).Should(Succeed())
			Expect(resp.Code).Should(Equal(200))
		})

		It("return error on FetchOrderByID", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(nil, errSomething)
			repo.EXPECT().FindRatingSellerBySellerID(gomock.Any()).Return(domain.SellerRating{}, errSomething)
			repo.EXPECT().CreateRatingSeller(gomock.Any()).Return(nil)
			_, err := orderUC.AddRatingSeller(mockRequest, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on FindRatingSellerBySellerID", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindRatingSellerBySellerID(gomock.Any()).Return(domain.SellerRating{}, errSomething)
			repo.EXPECT().CreateRatingSeller(gomock.Any()).Return(nil)
			_, err := orderUC.AddRatingSeller(mockRequest, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on CreateRatingSeller", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindRatingSellerBySellerID(gomock.Any()).Return(domain.SellerRating{}, nil)
			repo.EXPECT().CreateRatingSeller(gomock.Any()).Return(errSomething)
			_, err := orderUC.AddRatingSeller(mockRequest, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

		It("return error on UpdateRatingSeller", func() {
			repo.EXPECT().FetchOrderByID(gomock.Any()).Return(&mockOrder, nil)
			repo.EXPECT().FindRatingSellerBySellerID(gomock.Any()).Return(mockRating, nil)
			repo.EXPECT().UpdateRatingSeller(gomock.Any()).Return(errSomething)
			_, err := orderUC.AddRatingSeller(mockRequest, 1, 1)
			Expect(err).Should(HaveOccurred())
		})

	})

})
