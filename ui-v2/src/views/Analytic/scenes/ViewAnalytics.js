import React, { useState, useEffect } from 'react';
import {
  Table, Divider, Button, Popconfirm, Row, Col,
} from 'antd';
import { HitChartCard, HitCounterCard, UniquePageCounterCard } from 'components/Analytic';

function ViewAnalytics(props) {
  return (
    <div>
      <Row>
        <Col span={4} justify="space-around" align="middle">
          <HitCounterCard />
        </Col>
        <Col span={4} justify="space-around" align="middle">
          <UniquePageCounterCard />
        </Col>
      </Row>
      <Row>
        <Col span={24}>
          <HitChartCard />
        </Col>
      </Row>
    </div>
  );
}

export default ViewAnalytics;
