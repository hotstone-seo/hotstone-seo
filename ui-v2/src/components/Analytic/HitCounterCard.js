import React, { useState, useEffect, useCallback } from 'react';
import useInterval from '@use-hooks/interval';
import { fetchCountHit } from 'api/metric';
import CounterCard from './CounterCard';

function HitCounterCard({ ruleID }) {
  const [countHit, setCountHit] = useState(0);
  const [error, setError] = useState();

  const fetchData = useCallback(() => {
    const queryParm = ruleID ? { rule_id: ruleID } : {};
    fetchCountHit({ params: queryParm })
      .then((data) => {
        setCountHit(data.count);
      })
      .catch((error) => {
        setError(error);
      });
  }, [ruleID, setCountHit, setError]);

  useInterval(() => {
    fetchData();
  }, 5_000);

  useEffect(() => {
    fetchData();
  }, []);

  return <CounterCard counter={countHit} label="Hit" />;
}

export default HitCounterCard;
