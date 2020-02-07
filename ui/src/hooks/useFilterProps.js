import React, { useMemo, useRef } from "react";
import { Input, Button, Icon } from "antd";

export function useFilterProps(dataIndex) {
  const searchInputEl = useRef(null);

  const handleSearch = (selectedKeys, confirm) => {
    confirm();
  };

  const handleReset = clearFilters => {
    clearFilters();
  };

  return useMemo(() => {
    return {
      filterDropdown: ({
        setSelectedKeys,
        selectedKeys,
        confirm,
        clearFilters
      }) => (
        <div style={{ padding: 8 }}>
          <Input
            ref={searchInputEl}
            placeholder={`Search ${dataIndex}`}
            value={selectedKeys[0]}
            onChange={e =>
              setSelectedKeys(e.target.value ? [e.target.value] : [])
            }
            onPressEnter={() => handleSearch(selectedKeys, confirm, dataIndex)}
            style={{ width: 188, marginBottom: 8, display: "block" }}
          />
          <Button
            type="primary"
            onClick={() => handleSearch(selectedKeys, confirm, dataIndex)}
            icon="search"
            size="small"
            style={{ width: 90, marginRight: 8 }}
          >
            Search
          </Button>
          <Button
            onClick={() => handleReset(clearFilters)}
            size="small"
            style={{ width: 90 }}
          >
            Reset
          </Button>
        </div>
      ),

      filterIcon: filtered => (
        <Icon
          type="search"
          style={{ color: filtered ? "#1890ff" : undefined }}
        />
      ),

      onFilterDropdownVisibleChange: visible => {
        if (visible) {
          setTimeout(() => searchInputEl.current.select());
        }
      },

      render: text => {
        return text;
      }
    };
  }, [dataIndex]);
}
