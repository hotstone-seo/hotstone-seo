import React, { useState, useEffect, useCallback } from 'react';
import useInterval from '@use-hooks/interval';
import { fetchCountUniquePage } from 'api/metric';
import CounterCard from './CounterCard';

function UniquePageCounterCard({ ruleID }) {
  const [countUniquePage, setCountUniquePage] = useState(0);
  // const [setError] = useState();

  const fetchData = useCallback(() => {
    const queryParm = ruleID ? { rule_id: ruleID } : {};
    fetchCountUniquePage({ params: queryParm })
      .then((data) => {
        if (data !== undefined) setCountUniquePage(data.count);
      })
      .catch(() => {
        // handled in client.js
        // setError(error);
      });
  }, [ruleID, setCountUniquePage]);

  useInterval(() => {
    fetchData();
  }, 5_000);

  useEffect(() => {
    fetchData();
  }, [ruleID, fetchData]);

  return <CounterCard counter={countUniquePage} label="Unique Page" />;
}

export default UniquePageCounterCard;
