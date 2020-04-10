import React, {
  useState, useEffect, useReducer, useCallback,
} from 'react';
import {
  Card, Select, Typography, Alert,
} from 'antd';
import { ResponsiveLine } from '@nivo/line';
import useInterval from '@use-hooks/interval';
import { format, subDays, startOfMonth } from 'date-fns';
import { isAxiosError } from 'utils/axios';
import { fetchListCountHitPerDay } from 'api/metric';

const { Option } = Select;
const { Text } = Typography;

const FULL_DATE_FORMAT = 'yyyy-MM-dd';

const LAST_7_DAYS = 'last7Days';
const THIS_MONTH = 'thisMonth';

const initialState = {
  selectedOption: LAST_7_DAYS,
  startDate: format(subDays(new Date(), 7), FULL_DATE_FORMAT),
  endDate: format(new Date(), FULL_DATE_FORMAT),
};

function reducer(state, action) {
  const now = new Date();
  switch (action.type) {
    case LAST_7_DAYS:
      return {
        selectedOption: LAST_7_DAYS,
        startDate: format(subDays(now, 7), FULL_DATE_FORMAT),
        endDate: format(now, FULL_DATE_FORMAT),
      };
    case THIS_MONTH:
      return {
        selectedOption: THIS_MONTH,
        startDate: format(startOfMonth(now), FULL_DATE_FORMAT),
        endDate: format(now, FULL_DATE_FORMAT),
      };
    default:
      throw new Error();
  }
}

const rangeOptions = [
  { value: LAST_7_DAYS, label: 'Last 7 Days' },
  { value: THIS_MONTH, label: 'This Month' },
];

function toDataChart(dataListCountHit) {
  return dataListCountHit.map(({ date, count }) => ({ x: format(new Date(date), 'dd/MM'), y: count }));
}

function HitChartCard({ ruleID }) {
  const [state, dispatch] = useReducer(reducer, initialState);
  const [dataChart, setDataChart] = useState([{ id: 'count', data: [] }]);
  const [dataListCountHit, setDataListCountHit] = useState([]);
  const [error, setError] = useState();

  const fetchData = useCallback(
    (startDate, endDate, ruleID) => {
      let queryParm = { start: startDate, end: endDate };
      if (ruleID) {
        queryParm = { ...queryParm, rule_id: ruleID };
      }
      fetchListCountHitPerDay({ params: queryParm })
        .then((data) => {
          setDataListCountHit(data);
        })
        .catch((error) => {
          setError(error);
        });
    },
    [setDataListCountHit, setError],
  );

  // inspectAxiosError(error);
  const onChangeRange = (selectedOption) => {
    dispatch({ type: selectedOption });
  };

  useEffect(() => {
    if (dataListCountHit !== undefined) {
      setDataChart([{ id: 'count', data: toDataChart(dataListCountHit) }]);
    }
  }, [dataListCountHit]);

  useEffect(() => {
    fetchData(state.startDate, state.endDate, ruleID);
  }, [state, ruleID, fetchData]);

  useInterval(() => {
    fetchData(state.startDate, state.endDate, ruleID);
  }, 5_000);

  return (
    <Card>
      <div>
        {isAxiosError(error) && (
          <>
            <Alert type="error" message="Failed to fetch data because network error.Failed to connect API" />
            <br />
          </>
        )}

        <div>
          <Text strong>Hit Count in </Text>
          <Select value={state.selectedOption} onChange={onChangeRange}>
            {rangeOptions.map(({ value, label }, index) => (
              <Option key={value} value={value}>
                {label}
              </Option>
            ))}
          </Select>
        </div>
      </div>
      <div style={{ height: 400 }}>
        <ResponsiveLine
          data={dataChart}
          margin={{
            top: 50, right: 110, bottom: 50, left: 60,
          }}
          xScale={{
            type: 'point',
          }}
          yScale={{
            type: 'linear',
            min: 'auto',
            max: 'auto',
            stacked: true,
            reverse: false,
          }}
          axisBottom={{
            orient: 'bottom',
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: 'Date',
            legendOffset: 36,
            legendPosition: 'middle',
          }}
          axisLeft={{
            orient: 'left',
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: 'Hit Count',
            legendOffset: -40,
            legendPosition: 'middle',
          }}
          colors={{ scheme: 'nivo' }}
          pointSize={10}
          pointColor={{ theme: 'background' }}
          pointBorderWidth={2}
          pointBorderColor={{ from: 'serieColor' }}
          pointLabel="Hit Count"
          pointLabelYOffset={-12}
          enablePointLabel
          useMesh
        />
      </div>
    </Card>
  );
}

export default HitChartCard;
