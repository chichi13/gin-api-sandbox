package v1

import (
	"gin-api/schemas"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var fibonacciMemo map[int]float64

func GetFibonacci(c *gin.RouterGroup) {
	c.GET("/fibonacci/recursive/:n", func(c *gin.Context) {
		n, err := strconv.Atoi(c.Param("n"))
		if err != nil {
			// If the parameter is not an integer, return a bad request
			c.JSON(http.StatusBadRequest, schemas.BaseResponse{
				Success: false,
				Message: "Bad request",
			})
			return
		}
		c.JSON(http.StatusOK, schemas.BaseResponse{
			Success: true,
			Message: strconv.FormatFloat(recursiveFibonacci(n), 'f', 0, 64),
		})
	})
	c.GET("/fibonacci/iterative/:n", func(c *gin.Context) {
		n, err := strconv.Atoi(c.Param("n"))
		if err != nil {
			// If the parameter is not an integer, return a bad request
			c.JSON(http.StatusBadRequest, schemas.BaseResponse{
				Success: false,
				Message: "Bad request",
			})
			return
		}
		c.JSON(http.StatusOK, schemas.BaseResponse{
			Success: true,
			Message: strconv.FormatFloat(iterativeFibonacci(n), 'f', 0, 64),
		})
	})
}

// @Summary		Get the fibonacci number with recursive algorithm
// @Description	Get the fibonacci number with recursive algorithm
// @Accept			json
// @Produce		json
// @Param			n	path		int						true	"Number"
// @Success		200	{object}	schemas.BaseResponse	"Success"
// @Failure		400	{object}	schemas.BaseResponse	"Bad request"
// @Router			/fibonacci/recursive/{n} [get]
// @Tags			fibonacci
func recursiveFibonacci(n int) float64 {
	if n <= 1 {
		return float64(n)
	}

	if fibonacciMemo == nil {
		fibonacciMemo = map[int]float64{}
	}

	if val, ok := fibonacciMemo[n]; ok {
		return val
	}

	result := recursiveFibonacci(n-1) + recursiveFibonacci(n-2)
	fibonacciMemo[n] = result
	return result
}

// @Summary		Get the fibonacci number with iterative algorithm
// @Description	Get the fibonacci number with iterative algorithm
// @Accept			json
// @Produce		json
// @Param			n	path		int						true	"Number"
// @Success		200	{object}	schemas.BaseResponse	"Success"
// @Failure		400	{object}	schemas.BaseResponse	"Bad request"
// @Router			/fibonacci/iterative/{n} [get]
// @Tags			fibonacci
func iterativeFibonacci(n int) float64 {
	if n <= 1 {
		return float64(n)
	}

	var a, b, c float64
	a = 0
	b = 1
	for i := 2; i <= n; i++ {
		// a is the previous number fib(n-1)
		// b is the number before that fib(n-2)
		// c is the next number
		c = a + b
		a = b
		b = c
	}
	return c
}
