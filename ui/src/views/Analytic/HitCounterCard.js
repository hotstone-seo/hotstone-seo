import React, { useState, useEffect } from "react";
import CounterCard from "./CounterCard";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";
import useInterval from "@use-it/interval";

function HitCounterCard({ ruleID }) {
  const [countHit, setCountHit] = useState(0);
  const [{ data: dataCountHit }, refetch] = useHotstoneAPI({
    url: `metrics/hit?` + (ruleID ? `rule_id=${ruleID}` : ``)
  });

  useEffect(() => {
    if (dataCountHit !== undefined) {
      setCountHit(dataCountHit.count);
    }
  }, [dataCountHit]);

  useInterval(() => {
    refetch();
  }, 5_000);

  return <CounterCard counter={countHit} label="Hit" />;
}

export default HitCounterCard;
