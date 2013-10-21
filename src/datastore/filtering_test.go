package datastore

import (
	"common"
	. "launchpad.net/gocheck"
	"parser"
)

type FilteringSuite struct{}

var _ = Suite(&FilteringSuite{})

func (self *FilteringSuite) TestEqualityFiltering(c *C) {
	queryStr := "select * from t where column_one == 100 and column_two != 6;"
	query, err := parser.ParseQuery(queryStr)
	c.Assert(err, IsNil)
	series, err := common.StringToSeriesArray(`
[
 {
   "points": [
     {"values": [{"int64_value": 100},{"int64_value": 5 }], "timestamp": 1381346631, "sequence_number": 1},
     {"values": [{"int64_value": 100},{"int64_value": 6 }], "timestamp": 1381346631, "sequence_number": 1},
     {"values": [{"int64_value": 90 },{"int64_value": 15}], "timestamp": 1381346632, "sequence_number": 1}
   ],
   "name": "t",
   "fields": [
     {"type": "INT64", "name": "column_one"},
     {"type": "INT64", "name": "column_two"}
   ]
 }
]
`)
	c.Assert(err, IsNil)
	result, err := Filter(query, series[0])
	c.Assert(err, IsNil)
	c.Assert(result, NotNil)
	c.Assert(result.Points, HasLen, 1)
	c.Assert(*result.Points[0].Values[0].Int64Value, Equals, int64(100))
	c.Assert(*result.Points[0].Values[1].Int64Value, Equals, int64(5))
}

func (self *FilteringSuite) TestRegexFiltering(c *C) {
	queryStr := "select * from t where column_one ~= /.*foo.*/ and time > now() - 1d;"
	query, err := parser.ParseQuery(queryStr)
	c.Assert(err, IsNil)
	series, err := common.StringToSeriesArray(`
[
 {
   "points": [
     {"values": [{"string_value": "100"}], "timestamp": 1381346631, "sequence_number": 1},
     {"values": [{"string_value": "foobar"}], "timestamp": 1381346631, "sequence_number": 1}
   ],
   "name": "t",
   "fields": [
     {"type": "STRING", "name": "column_one"}
   ]
 }
]
`)
	c.Assert(err, IsNil)
	result, err := Filter(query, series[0])
	c.Assert(err, IsNil)
	c.Assert(result, NotNil)
	c.Assert(result.Points, HasLen, 1)
	c.Assert(*result.Points[0].Values[0].StringValue, Equals, "foobar")
}

func (self *FilteringSuite) TestInequalityFiltering(c *C) {
	queryStr := "select * from t where column_one >= 100 and column_two > 6 and time > now() - 1d;"
	query, err := parser.ParseQuery(queryStr)
	c.Assert(err, IsNil)
	series, err := common.StringToSeriesArray(`
[
 {
   "points": [
     {"values": [{"int64_value": 100},{"int64_value": 7 }], "timestamp": 1381346631, "sequence_number": 1},
     {"values": [{"int64_value": 100},{"int64_value": 6 }], "timestamp": 1381346631, "sequence_number": 1},
     {"values": [{"int64_value": 90 },{"int64_value": 15}], "timestamp": 1381346632, "sequence_number": 1}
   ],
   "name": "t",
   "fields": [
     {"type": "INT64", "name": "column_one"},
     {"type": "INT64", "name": "column_two"}
   ]
 }
]
`)
	c.Assert(err, IsNil)
	result, err := Filter(query, series[0])
	c.Assert(err, IsNil)
	c.Assert(result, NotNil)
	c.Assert(result.Points, HasLen, 1)
	c.Assert(*result.Points[0].Values[0].Int64Value, Equals, int64(100))
	c.Assert(*result.Points[0].Values[1].Int64Value, Equals, int64(7))
}