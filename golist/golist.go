package golist

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/goark/depm/subproc"
	"github.com/goark/errs"
)

type Context interface {
	GetPackagesRaw(context.Context, string) ([]byte, error)
	GetPackages(context.Context, string) ([]Package, error)
	GetModulesRaw(context.Context, string, bool) ([]byte, error)
	GetModules(context.Context, string, bool) ([]Module, error)
}

type ctx struct {
	conf *config
}

func New(opts ...OptEnv) Context {
	return &ctx{conf: newconfig(opts...)}
}

//GetPackagesRaw returns package information by JSON string
func (c *ctx) GetPackagesRaw(ctx context.Context, name string) ([]byte, error) {
	b, err := subproc.New("go", "list", "-json", name).
		WithContext(ctx).
		AddEnv(c.conf.GetEnv()...).
		WithStderr(c.conf.errorWriter).
		HideWindow(c.conf.flagHideWindow). //Windows only
		Output()
	return b, errs.Wrap(err, errs.WithContext("name", name))
}

//GetPackages returns package information
func (c *ctx) GetPackages(ctx context.Context, name string) ([]Package, error) {
	b, err := c.GetPackagesRaw(ctx, name)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("name", name))
	}
	decoder := json.NewDecoder(bytes.NewReader(b))
	ps := []Package{}
	for {
		var p Package
		if err := decoder.Decode(&p); err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return ps, errs.Wrap(err, errs.WithContext("name", name))
		}
		ps = append(ps, p)
	}
	return ps, nil
}

//GetModulesRaw returns module information by JSON string
func (c *ctx) GetModulesRaw(ctx context.Context, name string, updFlag bool) ([]byte, error) {
	args := []string{"list", "-json", "-m"}
	if updFlag {
		args = append(args, "-u")
	}
	args = append(args, name)
	b, err := subproc.New("go", args...).
		WithContext(ctx).
		AddEnv(c.conf.GetEnv()...).
		WithStderr(c.conf.errorWriter).
		HideWindow(true). //Windows only
		Output()
	return b, errs.Wrap(err, errs.WithContext("name", name))
}

//GetModules returns module information
func (c *ctx) GetModules(ctx context.Context, name string, updFlag bool) ([]Module, error) {
	b, err := c.GetModulesRaw(ctx, name, updFlag)
	if err != nil {
		return nil, errs.Wrap(err, errs.WithContext("name", name))
	}
	decoder := json.NewDecoder(bytes.NewReader(b))
	ms := []Module{}
	for {
		var m Module
		if err := decoder.Decode(&m); err != nil {
			if errs.Is(err, io.EOF) {
				break
			}
			return ms, errs.Wrap(err, errs.WithContext("name", name))
		}
		ms = append(ms, m)
	}
	return ms, nil
}

/* Copyright 2020 Spiegel
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * 	http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
