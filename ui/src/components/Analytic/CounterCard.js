import React from 'react';
import { Statistic, Card } from 'antd';

function CounterCard({ counter, label }) {
  return (
    <Card>
      <Statistic title={label} value={counter} />
    </Card>
  );
}

export default CounterCard;
