import { useEffect, useState } from 'react';
import { fetchDataSources } from 'api/datasource';

function useDataSources() {
  const [dataSources, setDataSources] = useState([]);

  useEffect(() => {
    fetchDataSources()
      .then((dataSources) => {
        if (dataSources !== undefined) setDataSources(dataSources);
      });
  }, []);

  return [dataSources, setDataSources];
}

export default useDataSources;
