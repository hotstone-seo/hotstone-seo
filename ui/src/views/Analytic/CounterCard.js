import React from "react";

function CounterCard({ counter, label }) {
  return (
    <div class="card text-center">
      <div class="card-body">
        <h5 class="card-text display-4">{counter}</h5>
      </div>
      <div class="card-footer">{label}</div>
    </div>
  );
}

export default CounterCard;
