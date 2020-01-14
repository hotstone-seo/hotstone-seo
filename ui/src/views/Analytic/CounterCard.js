import React from "react";

function CounterCard({ counter, label }) {
  return (
    <div className="card text-center">
      <div className="card-body">
        <h5 className="card-text display-4">{counter}</h5>
      </div>
      <div className="card-footer">{label}</div>
    </div>
  );
}

export default CounterCard;
