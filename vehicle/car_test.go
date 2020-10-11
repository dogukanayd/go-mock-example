package vehicle

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCARSTRUCT_Start(t *testing.T) {
	t.Run("it should return true when key equals to carkey and engine check return true", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		MockInterface := NewMockCAR(controller)
		var ctx *gin.Context

		MockInterface.EXPECT().EngineCheck(ctx, MockInterface).Return(true)
		result := CARSTRUCT{}.Start(ctx, "carkey", MockInterface)

		assert.Equal(t, true, result)
	})

	t.Run("it should return false when key not equals to carkey and engine check return true", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		MockInterface := NewMockCAR(controller)
		var ctx *gin.Context

		MockInterface.EXPECT().EngineCheck(ctx, MockInterface).Return(true)
		result := CARSTRUCT{}.Start(ctx, "qwe", MockInterface)

		assert.Equal(t, false, result)
	})

	t.Run("it should return false when engine check return false", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		MockInterface := NewMockCAR(controller)
		var ctx *gin.Context

		MockInterface.EXPECT().EngineCheck(ctx, MockInterface).Return(false)
		result := CARSTRUCT{}.Start(ctx, "carkey", MockInterface)

		assert.Equal(t, false, result)
	})
}

func TestCARSTRUCT_Start2(t *testing.T) {
	t.Run("it should return false", func(t *testing.T) {
		controller := gomock.NewController(t)
		defer controller.Finish()
		MockInterface := NewMockCAR(controller)
		var ctx *gin.Context
		result := CARSTRUCT{}.EngineCheck(ctx, MockInterface)

		assert.Equal(t, true, result)
	})
}
