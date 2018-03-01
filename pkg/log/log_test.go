// Copyright (c) 2017 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package log

// import (
// 	"context"
// 	"testing"
//
// 	"github.com/northwesternmutual/kanali/config"
// 	"github.com/spf13/viper"
// 	"github.com/stretchr/testify/assert"
// 	"go.uber.org/zap"
// 	"go.uber.org/zap/zapcore"
// 	"go.uber.org/zap/zaptest/observer"
// )
//
// func TestWithContext(t *testing.T) {
// 	logger := WithContext(nil)
// 	assert.Nil(t, logger)
//
// 	logger = WithContext(context.Background())
// 	assert.Nil(t, logger)
//
// 	viper.SetDefault(config.FlagProcessLogLevel.GetLong(), "foo")
// 	defer viper.Reset()
// 	Init(nil)
// 	logger = WithContext(nil)
// 	assert.True(t, logger.Core().Enabled(zapcore.InfoLevel))
// 	assert.False(t, logger.Core().Enabled(zapcore.DebugLevel))
//
// 	viper.SetDefault(config.FlagProcessLogLevel.GetLong(), "Warn")
// 	Init(nil)
// 	logger = WithContext(context.Background())
// 	assert.True(t, logger.Core().Enabled(zapcore.WarnLevel))
// 	assert.False(t, logger.Core().Enabled(zapcore.InfoLevel))
//
// 	core, obsvr := observer.New(zap.NewAtomicLevelAt(zapcore.InfoLevel))
// 	Init(core)
// 	assert.False(t, logger.Core().Enabled(zapcore.InfoLevel))
// 	assert.False(t, logger.Core().Enabled(zapcore.DebugLevel))
//
// 	ctx := NewContext(context.Background(), zap.String("foo", "bar"))
// 	logger = WithContext(ctx)
// 	logger.Info("test log")
// 	assert.Equal(t, 1, len(obsvr.All()[obsvr.Len()-1].Context))
// 	assert.Equal(t, "foo", obsvr.All()[obsvr.Len()-1].Context[0].Key)
// 	assert.Equal(t, "bar", obsvr.All()[obsvr.Len()-1].Context[0].String)
// }