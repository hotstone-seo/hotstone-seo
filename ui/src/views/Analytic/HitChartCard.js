import React, { useState, useEffect, useReducer } from "react";
import { ResponsiveLine } from "@nivo/line";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";
import useInterval from "@use-it/interval";
import { useForm } from "react-hook-form";
import { format, subDays, startOfMonth } from "date-fns";
import dataChart from "./data";
import inspectAxiosError, { isAxiosError } from "../../utils/axios";

const DATE_FORMAT = "yyyy-MM-dd";

const LAST_7_DAYS = "last7Days";
const THIS_MONTH = "thisMonth";

const initialState = {
  selectedOption: LAST_7_DAYS,
  startDate: format(subDays(new Date(), 7), DATE_FORMAT),
  endDate: format(new Date(), DATE_FORMAT)
};

function reducer(state, action) {
  const now = new Date();
  switch (action.type) {
    case LAST_7_DAYS:
      return {
        selectedOption: LAST_7_DAYS,
        startDate: format(subDays(now, 7), DATE_FORMAT),
        endDate: format(now, DATE_FORMAT)
      };
    case THIS_MONTH:
      return {
        selectedOption: THIS_MONTH,
        startDate: format(startOfMonth(now), DATE_FORMAT),
        endDate: format(now, DATE_FORMAT)
      };
    default:
      throw new Error();
  }
}

const rangeOptions = [
  { value: LAST_7_DAYS, label: "Last 7 Days" },
  { value: THIS_MONTH, label: "This Month" }
];

function HitChartCard() {
  const [state, dispatch] = useReducer(reducer, initialState);

  const { register, handleSubmit, errors } = useForm();

  const [countHit, setCountHit] = useState(0);
  const [{ data: dataListCountHit, error }, refetch] = useHotstoneAPI({
    url: `metrics/hit/range?start=${state.startDate}&end=${state.endDate}`
  });

  // inspectAxiosError(error);
  const onChangeRange = selectedOption => {
    console.log(selectedOption);
    dispatch({ type: selectedOption.range });
    refetch();
  };

  useEffect(() => {
    if (dataListCountHit !== undefined) {
      setCountHit(dataListCountHit.count);
    }
  }, [dataListCountHit]);

  useInterval(() => {
    refetch();
  }, 5_000);

  return (
    <div className="card">
      <div className="card-header text-left">
        {isAxiosError(error) && (
          <div class="alert alert-danger" role="alert">
            Failed to fetch data
          </div>
        )}
        <form>
          <select
            name="range"
            ref={register({ required: true })}
            onChange={handleSubmit(onChangeRange)}
          >
            {rangeOptions.map(({ value, label }, index) => {
              return (
                <option key={value} value={value}>
                  {label}
                </option>
              );
            })}
          </select>
        </form>
      </div>
      <div className="card-body" style={{ height: 400 }}>
        <ResponsiveLine
          data={dataChart}
          margin={{ top: 50, right: 110, bottom: 50, left: 60 }}
          xScale={{ type: "point" }}
          yScale={{
            type: "linear",
            min: "auto",
            max: "auto",
            stacked: true,
            reverse: false
          }}
          axisBottom={{
            orient: "bottom",
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: "Date",
            legendOffset: 36,
            legendPosition: "middle"
          }}
          axisLeft={{
            orient: "left",
            tickSize: 5,
            tickPadding: 5,
            tickRotation: 0,
            legend: "Hit Count",
            legendOffset: -40,
            legendPosition: "middle"
          }}
          colors={{ scheme: "nivo" }}
          pointSize={10}
          pointColor={{ theme: "background" }}
          pointBorderWidth={2}
          pointBorderColor={{ from: "serieColor" }}
          pointLabel="Hit Count"
          pointLabelYOffset={-12}
          useMesh={true}
        />
      </div>
    </div>
  );
}

export default HitChartCard;
