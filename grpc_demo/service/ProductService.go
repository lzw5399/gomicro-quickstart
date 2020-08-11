/**
 * @Author: lzw5399
 * @Date: 2020/8/11 22:31
 * @Desc: gRPC ProductService的实现类
 */
package service

import "context"

type ProdService struct {
}

func (ps *ProdService) GetProductStock(ctx context.Context, request *ProductRequest) (*ProductResponse, error) {
	return &ProductResponse{ProdStock: request.ProdId}, nil
}
