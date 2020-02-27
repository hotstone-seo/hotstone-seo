import React, { useState, useEffect, useCallback } from 'react';
import useInterval from '@use-hooks/interval';
import { fetchCountUniquePage } from 'api/metric';
import CounterCard from './CounterCard';

function UniquePageCounterCard({ ruleID }) {
  const [countUniquePage, setCountUniquePage] = useState(0);
  const [error, setError] = useState();

  const fetchData = useCallback(() => {
    const queryParm = ruleID ? { rule_id: ruleID } : {};
    fetchCountUniquePage({ params: queryParm })
      .then((data) => {
        setCountUniquePage(data.count);
      })
      .catch((error) => {
        setError(error);
      });
  }, [ruleID, setCountUniquePage, setError]);

  useInterval(() => {
    fetchData();
  }, 5_000);

  useEffect(() => {
    fetchData();
  }, []);

  return <CounterCard counter={countUniquePage} label="Unique Page" />;
}

export default UniquePageCounterCard;
