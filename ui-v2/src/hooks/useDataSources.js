import { useEffect, useState } from 'react';
import { fetchDataSources } from 'api/datasource';

function useDataSources() {
  const [dataSources, setDataSources] = useState([]);

  useEffect(() => {
    fetchDataSources()
      .then((dataSources) => {
        setDataSources(dataSources);
      });
  }, []);

  return dataSources;
}

export default useDataSources;
