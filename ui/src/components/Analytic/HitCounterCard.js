import React, { useState, useEffect, useCallback } from 'react';
import useInterval from '@use-hooks/interval';
import { fetchCountHit } from 'api/metric';
import CounterCard from './CounterCard';

function HitCounterCard({ ruleID }) {
  const [countHit, setCountHit] = useState(0);

  const fetchData = useCallback(() => {
    const queryParm = ruleID ? { rule_id: ruleID } : {};
    fetchCountHit({ params: queryParm })
      .then((data) => {
        if (data !== undefined) setCountHit(data.count);
      })
      .catch(() => {
        // handle in client.js
        // setError(error);
      });
  }, [ruleID]);

  useInterval(() => {
    fetchData();
  }, 5_000);

  useEffect(() => {
    fetchData();
  }, [ruleID, fetchData]);

  return <CounterCard counter={countHit} label="Hit" />;
}

export default HitCounterCard;
