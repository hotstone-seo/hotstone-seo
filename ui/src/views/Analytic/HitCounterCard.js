import React, { useState, useEffect } from "react";
import CounterCard from "./CounterCard";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";

function HitCounterCard() {
  const [countHit, setCountHit] = useState(0);
  const { data: dataCountHit, loading, timer } = useHotstoneAPI({
    url: "metrics/hit",
    pollingInterval: 5000
  });
  useEffect(() => {
    if (dataCountHit !== undefined) {
      setCountHit(dataCountHit.count);
    }
  }, [dataCountHit]);

  return <CounterCard counter={countHit} label="Hit" />;
}

export default HitCounterCard;
