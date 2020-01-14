import React from "react";
import CounterCard from "./CounterCard";
import { ResponsiveLine } from "@nivo/line";
import { useForm } from "react-hook-form";

import data from "./data";

function AnalyticPage() {
  const { register, handleSubmit, errors } = useForm();
  const onChangeRange = data => console.log(data);
  console.log(errors);

  return (
    <div className="container">
      <div className="row">
        <div className="col">
          <CounterCard counter={234} label="Hit" />
        </div>
        <div className="col">
          <CounterCard counter={45} label="Unique Page" />
        </div>
        <div className="col"></div>
      </div>
      <div className="row">
        <div className="col">
          <div className="card" style={{ height: 400 }}>
            <div className="card-header text-left">
              <form>
                <select
                  name="range"
                  ref={register({ required: true })}
                  onChange={handleSubmit(onChangeRange)}
                >
                  <option value="last-7days">Last 7 Days</option>
                  <option value="last-3days">Last 3 Days</option>
                </select>
              </form>
            </div>
            <div className="card-body">
              <ResponsiveLine
                data={data}
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
        </div>
      </div>
    </div>
  );
}

export default AnalyticPage;

// TODO: Big question: How to combine multiple reducers (and stores) to compose a global store ? i.e. RootStore = AuthStore + LocaleStore + ...
// TODO: Have to use new way of React: useState, useReducer, useContext
// Bookmarks:
// - https://leewarrick.com/blog/a-guide-to-usestate-and-usereducer/
// - https://medium.com/crowdbotics/how-to-use-usereducer-in-react-hooks-for-performance-optimization-ecafca9e7bf5
// - https://dev.to/ramsay/build-a-redux-like-global-store-using-react-hooks-4a7n
