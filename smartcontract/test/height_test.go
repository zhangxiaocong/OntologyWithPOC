/*
 * Copyright (C) 2018 The ontology Authors
 * This file is part of The ontology library.
 *
 * The ontology is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Lesser General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * The ontology is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU Lesser General Public License for more details.
 *
 * You should have received a copy of the GNU Lesser General Public License
 * along with The ontology.  If not, see <http://www.gnu.org/licenses/>.
 */

package test

import (
	"OntologyWithPOC/smartcontract"
	"OntologyWithPOC/vm/neovm"
	"OntologyWithPOC/vm/neovm/errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeight(t *testing.T) {
	byteCode0 := []byte{
		byte(neovm.NEWMAP),
		byte(neovm.PUSH0),
		byte(neovm.HASKEY),
	}

	byteCode1 := []byte{
		byte(neovm.NEWMAP),
		byte(neovm.KEYS),
	}

	byteCode2 := []byte{
		byte(neovm.NEWMAP),
		byte(neovm.VALUES),
	}

	bytecode := [...][]byte{byteCode0, byteCode1, byteCode2}

	for i := 0; i < 3; i++ {
		config := &smartcontract.Config{
			Time:   10,
			Height: 10,
			//Tx:     &types.Transaction{},
		}
		sc := smartcontract.SmartContract{
			Config:  config,
			Gas:     100,
			CacheDB: nil,
		}
		engine, err := sc.NewExecuteEngine(bytecode[i])

		_, err = engine.Invoke()

		assert.EqualError(t, err, "[NeoVmService] vm execution error!: "+errors.ERR_NOT_SUPPORT_OPCODE.Error())
	}
}
