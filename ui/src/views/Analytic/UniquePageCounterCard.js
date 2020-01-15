import React, { useState, useEffect } from "react";
import CounterCard from "./CounterCard";
import useHotstoneAPI from "../../hooks/useHotstoneAPI";

function UniquePageCounterCard() {
  const [countUniquePage, setCountUniquePage] = useState(0);
  const { data: dataCountUniquePage, error, params } = useHotstoneAPI({
    url: "metrics/unique-page",
    pollingInterval: 5000
  });
  console.log("ERROR unique page: " + error);
  console.log("PARAMS unique page: " + params);
  useEffect(() => {
    if (dataCountUniquePage !== undefined) {
      setCountUniquePage(dataCountUniquePage.count);
    }
  }, [dataCountUniquePage]);

  return <CounterCard counter={countUniquePage} label="Unique Page" />;
}

export default UniquePageCounterCard;
