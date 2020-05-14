import React from 'react';
import { Card, Result } from 'antd';

function GenericNotAuthorized() {
  return (
    <div className="GenericNotAuthorized">
      <Card>
        <Result
          status="403"
          title="403"
          subTitle="Sorry, you are not authorized to access this page."
        />
      </Card>
    </div>
  );
}

export default GenericNotAuthorized;