import React, { useState, useEffect } from "react";
import CounterCard from "./CounterCard";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";
import useInterval from "@use-it/interval";

function UniquePageCounterCard({ ruleID }) {
  const [countUniquePage, setCountUniquePage] = useState(0);
  const [{ data: dataCountUniquePage }, refetch] = useHotstoneAPI({
    url: "metrics/unique-page?" + (ruleID ? `rule_id=${ruleID}` : ``)
  });

  useEffect(() => {
    if (dataCountUniquePage !== undefined) {
      setCountUniquePage(dataCountUniquePage.count);
    }
  }, [dataCountUniquePage]);

  useInterval(() => {
    refetch();
  }, 5_000);

  return <CounterCard counter={countUniquePage} label="Unique Page" />;
}

export default UniquePageCounterCard;
