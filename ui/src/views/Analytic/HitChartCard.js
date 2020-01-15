import React, { useState, useEffect } from "react";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";
import useInterval from "@use-it/interval";

const LAST_7_DAYS = "last7Days";
const THIS_MONTH = "thisMonth";

const initialState = {
  selectedOption: "",
  startDate: null,
  endDate: null
};

const rangeOptions = [
  { value: LAST_7_DAYS, label: "Last 7 Days" },
  { value: THIS_MONTH, label: "This Month" }
];

function HitChartCard() {
  const { register, handleSubmit, errors } = useForm();
  const onChangeRange = data => console.log(data);

  const [countHit, setCountHit] = useState(0);
  const [{ data: dataCountHit }, refetch] = useHotstoneAPI({
    url: "metrics/hit"
  });

  useEffect(() => {
    if (dataCountHit !== undefined) {
      setCountHit(dataCountHit.count);
    }
  }, [dataCountHit]);

  useInterval(() => {
    refetch();
  }, 5_000);

  return (
    <div className="card" style={{ height: 400 }}>
      <div className="card-header text-left">
        <form>
          <select
            name="range"
            ref={register({ required: true })}
            onChange={handleSubmit(onChangeRange)}
          >
            {rangeOptions.map(({ value, label }, index) => {
              return <option value={value}>{label}</option>;
            })}
          </select>
        </form>
      </div>
      <div className="card-body">
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
            legend: "transportation",
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
