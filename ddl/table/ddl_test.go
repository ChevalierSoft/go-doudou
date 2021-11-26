package table

import (
	"context"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/unionj-cloud/go-doudou/ddl/columnenum"
	"github.com/unionj-cloud/go-doudou/ddl/sortenum"
	"github.com/unionj-cloud/go-doudou/ddl/wrapper"
	"reflect"
	"testing"
)

func ExampleCreateTable() {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	expectjson := `{"Name":"user_createtable","Columns":[{"Table":"user","Name":"id","Type":"INT","Default":null,"Pk":true,"Nullable":false,"Unsigned":false,"Autoincrement":true,"Extra":"","Meta":{"Name":"ID","Type":"int","Tag":"dd:\"pk;auto\"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"name","Type":"VARCHAR(255)","Default":"'jack'","Pk":false,"Nullable":false,"Unsigned":false,"Autoincrement":false,"Extra":"","Meta":{"Name":"Name","Type":"string","Tag":"dd:\"index:name_phone_idx,2;default:'jack'\"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"phone","Type":"VARCHAR(255)","Default":"'13552053960'","Pk":false,"Nullable":false,"Unsigned":false,"Autoincrement":false,"Extra":"comment '手机号'","Meta":{"Name":"Phone","Type":"string","Tag":"dd:\"index:name_phone_idx,1;default:'13552053960';extra:comment '手机号'\"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"age","Type":"INT","Default":null,"Pk":false,"Nullable":false,"Unsigned":false,"Autoincrement":false,"Extra":"","Meta":{"Name":"Age","Type":"int","Tag":"dd:\"index\"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"no","Type":"INT","Default":null,"Pk":false,"Nullable":false,"Unsigned":false,"Autoincrement":false,"Extra":"","Meta":{"Name":"No","Type":"int","Tag":"dd:\"unique\"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"school","Type":"VARCHAR(255)","Default":"'harvard'","Pk":false,"Nullable":true,"Unsigned":false,"Autoincrement":false,"Extra":"comment '学校'","Meta":{"Name":"School","Type":"string","Tag":"dd:\"null;default:'harvard';extra:comment '学校'\"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"is_student","Type":"TINYINT","Default":null,"Pk":false,"Nullable":false,"Unsigned":false,"Autoincrement":false,"Extra":"","Meta":{"Name":"IsStudent","Type":"bool","Tag":"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"delete_at","Type":"DATETIME","Default":null,"Pk":false,"Nullable":true,"Unsigned":false,"Autoincrement":false,"Extra":"","Meta":{"Name":"DeleteAt","Type":"*time.Time","Tag":"","Comments":null},"AutoSet":false,"Indexes":null},{"Table":"user","Name":"create_at","Type":"DATETIME","Default":"CURRENT_TIMESTAMP","Pk":false,"Nullable":true,"Unsigned":false,"Autoincrement":false,"Extra":"","Meta":{"Name":"CreateAt","Type":"*time.Time","Tag":"dd:\"default:CURRENT_TIMESTAMP\"","Comments":null},"AutoSet":true,"Indexes":null},{"Table":"user","Name":"update_at","Type":"DATETIME","Default":"CURRENT_TIMESTAMP","Pk":false,"Nullable":true,"Unsigned":false,"Autoincrement":false,"Extra":"ON UPDATE CURRENT_TIMESTAMP","Meta":{"Name":"UpdateAt","Type":"*time.Time","Tag":"dd:\"default:CURRENT_TIMESTAMP;extra:ON UPDATE CURRENT_TIMESTAMP\"","Comments":null},"AutoSet":true,"Indexes":null}],"Pk":"id","Indexes":[{"Unique":false,"Name":"name_phone_idx","Items":[{"Unique":false,"Name":"","Column":"phone","Order":1,"Sort":"asc"},{"Unique":false,"Name":"","Column":"name","Order":2,"Sort":"asc"}]},{"Unique":false,"Name":"age_idx","Items":[{"Unique":false,"Name":"","Column":"age","Order":1,"Sort":"asc"}]},{"Unique":true,"Name":"no_idx","Items":[{"Unique":false,"Name":"","Column":"no","Order":1,"Sort":"asc"}]}],"Meta":{"Name":"User","Fields":[{"Name":"ID","Type":"int","Tag":"dd:\"pk;auto\"","Comments":null},{"Name":"Name","Type":"string","Tag":"dd:\"index:name_phone_idx,2;default:'jack'\"","Comments":null},{"Name":"Phone","Type":"string","Tag":"dd:\"index:name_phone_idx,1;default:'13552053960';extra:comment '手机号'\"","Comments":null},{"Name":"Age","Type":"int","Tag":"dd:\"index\"","Comments":null},{"Name":"No","Type":"int","Tag":"dd:\"unique\"","Comments":null},{"Name":"School","Type":"string","Tag":"dd:\"null;default:'harvard';extra:comment '学校'\"","Comments":null},{"Name":"IsStudent","Type":"bool","Tag":"","Comments":null},{"Name":"DeleteAt","Type":"*time.Time","Tag":"","Comments":null},{"Name":"CreateAt","Type":"*time.Time","Tag":"dd:\"default:CURRENT_TIMESTAMP\"","Comments":null},{"Name":"UpdateAt","Type":"*time.Time","Tag":"dd:\"default:CURRENT_TIMESTAMP;extra:ON UPDATE CURRENT_TIMESTAMP\"","Comments":null}],"Comments":["dd:table"],"Methods":null}}`
	var table Table
	if err = json.Unmarshal([]byte(expectjson), &table); err != nil {
		panic(err)
	}
	if err := CreateTable(context.Background(), db, table); (err != nil) != false {
		panic(fmt.Sprintf("CreateTable() error = %v, wantErr %v", err, false))
	}

	// Output:

}

func TestChangeColumn(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	type args struct {
		db  *sqlx.DB
		col Column
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errmsg  string
	}{
		{
			name: "1",
			args: args{
				db: db,
				col: Column{
					Table:   "ddl_user",
					Name:    "school",
					Type:    "varchar(45)",
					Default: "'Beijing Univ.'",
				},
			},
			wantErr: false,
		},
		{
			name: "2",
			args: args{
				db: db,
				col: Column{
					Table:   "ddl_user",
					Name:    "school",
					Type:    columnenum.TextType,
					Default: "'Beijing Univ.'",
				},
			},
			wantErr: true,
			errmsg:  `Error 1101: BLOB, TEXT, GEOMETRY or JSON column 'school' can't have a default value`,
		},
		{
			name: "3",
			args: args{
				db: db,
				col: Column{
					Table:   "ddl_user",
					Name:    "school",
					Type:    "varchar(45)",
					Default: "Beijing Univ.",
				},
			},
			wantErr: true,
			errmsg:  `Error 1064: You have an error in your SQL syntax; check the manual that corresponds to your MySQL server version for the right syntax to use near 'Beijing Univ.' at line 2`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if err = ChangeColumn(context.Background(), tt.args.db, tt.args.col); (err != nil) != tt.wantErr {
				t.Errorf("ChangeColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil {
				if err.Error() != tt.errmsg {
					t.Errorf("want %s, got %s", tt.errmsg, err.Error())
				}
			}
		})
	}
}

func TestAddColumn(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	type args struct {
		db  *sqlx.DB
		col Column
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		errmsg  string
	}{
		{
			name: "1",
			args: args{
				db: db,
				col: Column{
					Table:   "ddl_user",
					Name:    "favourite",
					Type:    "varchar(45)",
					Default: "'football'",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var err error
			if err = AddColumn(context.Background(), tt.args.db, tt.args.col); (err != nil) != tt.wantErr {
				t.Errorf("ChangeColumn() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDropIndex(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	type args struct {
		ctx context.Context
		db  wrapper.Querier
		idx Index
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				db:  db,
				idx: Index{
					Table:  "ddl_user",
					Unique: true,
					Name:   "age_idx",
					Items: []IndexItem{
						{
							Column: "age",
							Order:  1,
							Sort:   sortenum.Asc,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DropIndex(tt.args.ctx, tt.args.db, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("DropIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAddIndex(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	type args struct {
		ctx context.Context
		db  wrapper.Querier
		idx Index
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				db:  db,
				idx: Index{
					Table:  "ddl_user",
					Unique: true,
					Name:   "school_idx",
					Items: []IndexItem{
						{
							Column: "school",
							Order:  1,
							Sort:   sortenum.Asc,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddIndex(tt.args.ctx, tt.args.db, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("AddIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDropAddIndex(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	type args struct {
		ctx context.Context
		db  wrapper.Querier
		idx Index
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "",
			args: args{
				ctx: context.Background(),
				db:  db,
				idx: Index{
					Table:  "ddl_user",
					Unique: true,
					Name:   "age_idx",
					Items: []IndexItem{
						{
							Column: "age",
							Order:  1,
							Sort:   sortenum.Asc,
						},
						{
							Column: "school",
							Order:  2,
							Sort:   sortenum.Asc,
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := DropAddIndex(tt.args.ctx, tt.args.db, tt.args.idx); (err != nil) != tt.wantErr {
				t.Errorf("DropAddIndex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_foreignKeys(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	type args struct {
		ctx    context.Context
		db     *sqlx.DB
		schema string
		t      string
	}
	tests := []struct {
		name    string
		args    args
		wantFks []ForeignKey
	}{
		{
			name: "",
			args: args{
				ctx:    context.Background(),
				db:     db,
				schema: "test",
				t:      "ddl_book",
			},
			wantFks: []ForeignKey{
				{
					Table:           "ddl_book",
					Constraint:      "fk_user",
					Fk:              "user_id",
					ReferencedTable: "ddl_user",
					ReferencedCol:   "id",
					UpdateRule:      "NO ACTION",
					DeleteRule:      "NO ACTION",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotFks := foreignKeys(tt.args.ctx, tt.args.db, tt.args.schema, tt.args.t); !reflect.DeepEqual(gotFks, tt.wantFks) {
				t.Errorf("foreignKeys() = %v, want %v", gotFks, tt.wantFks)
			}
		})
	}
}

func TestTable2struct(t *testing.T) {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()
	expectJson := `[
  {
    "Name": "ddl_book",
    "Columns": [
      {
        "Table": "ddl_book",
        "Name": "id",
        "Type": "int",
        "Default": "",
        "Pk": true,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": true,
        "Extra": "",
        "Meta": {
          "Name": "Id",
          "Type": "int",
          "Tag": "dd:\"pk;auto;type:int\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": true,
            "Name": "PRIMARY",
            "Column": "id",
            "Order": 1,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_book",
        "Name": "name",
        "Type": "varchar(45)",
        "Default": "",
        "Pk": false,
        "Nullable": true,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "Name",
          "Type": "*string",
          "Tag": "dd:\"type:varchar(45)\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": null,
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_book",
        "Name": "user_id",
        "Type": "int",
        "Default": "",
        "Pk": false,
        "Nullable": true,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "UserId",
          "Type": "*int",
          "Tag": "dd:\"type:int;index:fk_user,1,asc;fk:ddl_user,id,fk_user,ON DELETE NO ACTION ON UPDATE NO ACTION\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": false,
            "Name": "fk_user",
            "Column": "user_id",
            "Order": 1,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "ddl_book",
          "Constraint": "fk_user",
          "Fk": "user_id",
          "ReferencedTable": "ddl_user",
          "ReferencedCol": "id",
          "UpdateRule": "NO ACTION",
          "DeleteRule": "NO ACTION"
        }
      }
    ],
    "Pk": "id",
    "Indexes": [
      {
        "Table": "",
        "Unique": true,
        "Name": "PRIMARY",
        "Items": [
          {
            "Unique": true,
            "Name": "PRIMARY",
            "Column": "id",
            "Order": 1,
            "Sort": "asc"
          }
        ]
      },
      {
        "Table": "",
        "Unique": false,
        "Name": "fk_user",
        "Items": [
          {
            "Unique": false,
            "Name": "fk_user",
            "Column": "user_id",
            "Order": 1,
            "Sort": "asc"
          }
        ]
      }
    ],
    "Meta": {
      "Name": "Book",
      "Fields": [
        {
          "Name": "Id",
          "Type": "int",
          "Tag": "dd:\"pk;auto;type:int\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "Name",
          "Type": "*string",
          "Tag": "dd:\"type:varchar(45)\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "UserId",
          "Type": "*int",
          "Tag": "dd:\"type:int;index:fk_user,1,asc;fk:ddl_user,id,fk_user,ON DELETE NO ACTION ON UPDATE NO ACTION\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        }
      ],
      "Comments": null,
      "Methods": null,
      "IsExport": false
    },
    "Fks": [
      {
        "Table": "ddl_book",
        "Constraint": "fk_user",
        "Fk": "user_id",
        "ReferencedTable": "ddl_user",
        "ReferencedCol": "id",
        "UpdateRule": "NO ACTION",
        "DeleteRule": "NO ACTION"
      }
    ]
  },
  {
    "Name": "ddl_user",
    "Columns": [
      {
        "Table": "ddl_user",
        "Name": "id",
        "Type": "int",
        "Default": "",
        "Pk": true,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": true,
        "Extra": "",
        "Meta": {
          "Name": "Id",
          "Type": "int",
          "Tag": "dd:\"pk;auto;type:int\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": true,
            "Name": "PRIMARY",
            "Column": "id",
            "Order": 1,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "name",
        "Type": "varchar(255)",
        "Default": "jack",
        "Pk": false,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "Name",
          "Type": "string",
          "Tag": "dd:\"type:varchar(255);default:'jack';index:name_phone_idx,2,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": false,
            "Name": "name_phone_idx",
            "Column": "name",
            "Order": 2,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "phone",
        "Type": "varchar(255)",
        "Default": "13552053960",
        "Pk": false,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "comment 'mobile phone'",
        "Meta": {
          "Name": "Phone",
          "Type": "string",
          "Tag": "dd:\"type:varchar(255);default:'13552053960';extra:comment 'mobile phone';index:name_phone_idx,1,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": false,
            "Name": "name_phone_idx",
            "Column": "phone",
            "Order": 1,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "age",
        "Type": "int",
        "Default": "",
        "Pk": false,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "Age",
          "Type": "int",
          "Tag": "dd:\"type:int;index:age_idx,1,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": false,
            "Name": "age_idx",
            "Column": "age",
            "Order": 1,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "no",
        "Type": "int",
        "Default": "",
        "Pk": false,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "No",
          "Type": "int",
          "Tag": "dd:\"type:int;unique:no_idx,1,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": [
          {
            "Unique": true,
            "Name": "no_idx",
            "Column": "no",
            "Order": 1,
            "Sort": "asc"
          }
        ],
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "school",
        "Type": "varchar(255)",
        "Default": "harvard",
        "Pk": false,
        "Nullable": true,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "comment 'school'",
        "Meta": {
          "Name": "School",
          "Type": "*string",
          "Tag": "dd:\"type:varchar(255);default:'harvard';extra:comment 'school'\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": null,
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "is_student",
        "Type": "tinyint",
        "Default": "",
        "Pk": false,
        "Nullable": false,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "IsStudent",
          "Type": "int8",
          "Tag": "dd:\"type:tinyint\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": null,
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "create_at",
        "Type": "datetime",
        "Default": "CURRENT_TIMESTAMP",
        "Pk": false,
        "Nullable": true,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "CreateAt",
          "Type": "*time.Time",
          "Tag": "dd:\"type:datetime;default:CURRENT_TIMESTAMP\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": true,
        "Indexes": null,
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "update_at",
        "Type": "datetime",
        "Default": "CURRENT_TIMESTAMP",
        "Pk": false,
        "Nullable": true,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "on update CURRENT_TIMESTAMP",
        "Meta": {
          "Name": "UpdateAt",
          "Type": "*time.Time",
          "Tag": "dd:\"type:datetime;default:CURRENT_TIMESTAMP;extra:on update CURRENT_TIMESTAMP\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": true,
        "Indexes": null,
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      },
      {
        "Table": "ddl_user",
        "Name": "delete_at",
        "Type": "datetime",
        "Default": "",
        "Pk": false,
        "Nullable": true,
        "Unsigned": false,
        "Autoincrement": false,
        "Extra": "",
        "Meta": {
          "Name": "DeleteAt",
          "Type": "*time.Time",
          "Tag": "dd:\"type:datetime\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        "AutoSet": false,
        "Indexes": null,
        "Fk": {
          "Table": "",
          "Constraint": "",
          "Fk": "",
          "ReferencedTable": "",
          "ReferencedCol": "",
          "UpdateRule": "",
          "DeleteRule": ""
        }
      }
    ],
    "Pk": "id",
    "Indexes": [
      {
        "Table": "",
        "Unique": true,
        "Name": "PRIMARY",
        "Items": [
          {
            "Unique": true,
            "Name": "PRIMARY",
            "Column": "id",
            "Order": 1,
            "Sort": "asc"
          }
        ]
      },
      {
        "Table": "",
        "Unique": true,
        "Name": "no_idx",
        "Items": [
          {
            "Unique": true,
            "Name": "no_idx",
            "Column": "no",
            "Order": 1,
            "Sort": "asc"
          }
        ]
      },
      {
        "Table": "",
        "Unique": false,
        "Name": "age_idx",
        "Items": [
          {
            "Unique": false,
            "Name": "age_idx",
            "Column": "age",
            "Order": 1,
            "Sort": "asc"
          }
        ]
      },
      {
        "Table": "",
        "Unique": false,
        "Name": "name_phone_idx",
        "Items": [
          {
            "Unique": false,
            "Name": "name_phone_idx",
            "Column": "phone",
            "Order": 1,
            "Sort": "asc"
          },
          {
            "Unique": false,
            "Name": "name_phone_idx",
            "Column": "name",
            "Order": 2,
            "Sort": "asc"
          }
        ]
      }
    ],
    "Meta": {
      "Name": "User",
      "Fields": [
        {
          "Name": "Id",
          "Type": "int",
          "Tag": "dd:\"pk;auto;type:int\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "Name",
          "Type": "string",
          "Tag": "dd:\"type:varchar(255);default:'jack';index:name_phone_idx,2,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "Phone",
          "Type": "string",
          "Tag": "dd:\"type:varchar(255);default:'13552053960';extra:comment 'mobile phone';index:name_phone_idx,1,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "Age",
          "Type": "int",
          "Tag": "dd:\"type:int;index:age_idx,1,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "No",
          "Type": "int",
          "Tag": "dd:\"type:int;unique:no_idx,1,asc\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "School",
          "Type": "*string",
          "Tag": "dd:\"type:varchar(255);default:'harvard';extra:comment 'school'\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "IsStudent",
          "Type": "int8",
          "Tag": "dd:\"type:tinyint\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "CreateAt",
          "Type": "*time.Time",
          "Tag": "dd:\"type:datetime;default:CURRENT_TIMESTAMP\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "UpdateAt",
          "Type": "*time.Time",
          "Tag": "dd:\"type:datetime;default:CURRENT_TIMESTAMP;extra:on update CURRENT_TIMESTAMP\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        },
        {
          "Name": "DeleteAt",
          "Type": "*time.Time",
          "Tag": "dd:\"type:datetime\"",
          "Comments": null,
          "IsExport": false,
          "DocName": ""
        }
      ],
      "Comments": null,
      "Methods": null,
      "IsExport": false
    },
    "Fks": null
  }
]
`
	var wantTables []Table
	_ = json.Unmarshal([]byte(expectJson), &wantTables)

	type args struct {
		ctx         context.Context
		dir         string
		pre         string
		schema      string
		existTables []string
		db          *sqlx.DB
	}
	tests := []struct {
		name       string
		args       args
		wantTables []Table
	}{
		{
			name: "",
			args: args{
				ctx:         context.Background(),
				pre:         "ddl_",
				schema:      "test",
				existTables: []string{"ddl_book", "ddl_user"},
				db:          db,
			},
			wantTables: wantTables,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotTables := Table2struct(tt.args.ctx, tt.args.pre, tt.args.schema, tt.args.existTables, tt.args.db)
			assert.Equal(t, len(gotTables), len(tt.wantTables))
			var nameEqualCount int
			for _, item := range gotTables {
				for _, _item := range tt.wantTables {
					if _item.Name == item.Name {
						nameEqualCount++
						assert.ElementsMatch(t, item.Indexes, _item.Indexes)
						assert.ElementsMatch(t, item.Fks, _item.Fks)
						assert.ElementsMatch(t, item.Columns, _item.Columns)
						break
					}
				}
			}
			assert.Equal(t, nameEqualCount, len(gotTables))
		})
	}
}

func ExampleStruct2Table() {
	terminator, db, err := Setup()
	if err != nil {
		panic(err)
	}
	defer terminator()
	defer db.Close()

	_ = Struct2Table(context.Background(), "../testdata/domain", "ddl_", []string{"ddl_user", "ddl_book"}, db)
	// Output:
	//CREATE TABLE `ddl_order` (
	//`id` INT NOT NULL AUTO_INCREMENT,
	//`amount` BIGINT NOT NULL,
	//`user_id` int NOT NULL,
	//`create_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP,
	//`delete_at` DATETIME NULL,
	//`update_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	//PRIMARY KEY (`id`),
	//CONSTRAINT `fk_ddl_user` FOREIGN KEY (`user_id`)
	//REFERENCES `ddl_user`(`id`)
	//ON DELETE CASCADE ON UPDATE NO ACTION)
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `id` `id` INT NOT NULL AUTO_INCREMENT;
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `name` `name` VARCHAR(255) NOT NULL DEFAULT 'jack';
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `phone` `phone` VARCHAR(255) NOT NULL DEFAULT '13552053960' comment '手机号';
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `age` `age` INT NOT NULL;
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `no` `no` int NOT NULL;
	//ALTER TABLE `ddl_user`
	//ADD COLUMN `unique_col` int NOT NULL;
	//ALTER TABLE `ddl_user`
	//ADD COLUMN `unique_col_2` int NOT NULL;
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `school` `school` VARCHAR(255) NULL DEFAULT 'harvard' comment '学校';
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `is_student` `is_student` TINYINT NOT NULL;
	//ALTER TABLE `ddl_user`
	//ADD COLUMN `rule` varchar(255) NOT NULL comment '链接匹配规则，匹配的链接采用该css规则来爬';
	//ALTER TABLE `ddl_user`
	//ADD COLUMN `rule_type` varchar(45) NOT NULL comment '链接匹配规则类型，支持prefix前缀匹配和regex正则匹配';
	//ALTER TABLE `ddl_user`
	//ADD COLUMN `arrive_at` datetime NULL comment '到货时间';
	//ALTER TABLE `ddl_user`
	//ADD COLUMN `status` tinyint(4) NOT NULL comment '0进行中
	//1完结
	//2取消';
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `create_at` `create_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP;
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `delete_at` `delete_at` DATETIME NULL;
	//ALTER TABLE `ddl_user`
	//CHANGE COLUMN `update_at` `update_at` DATETIME NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;
	//ALTER TABLE `ddl_user` ADD UNIQUE INDEX `rule_idx` (`rule` asc);
	//ALTER TABLE `ddl_user` ADD UNIQUE INDEX `unique_col_idx` (`unique_col` asc,`unique_col_2` asc);
}
