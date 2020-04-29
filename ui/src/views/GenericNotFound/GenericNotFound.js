import React from 'react';
import { Card, Alert, Result, } from 'antd';

function GenericNotFound() {
  return (
    <div className="GenericNotFound">
      <Card>
        <Result
          status="404"
          title="Page Not Found"
          subTitle="Sorry, we could't process your request page"
        />
      </Card>
    </div>
  );
}

export default GenericNotFound;
