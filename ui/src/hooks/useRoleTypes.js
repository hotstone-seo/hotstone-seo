import { useEffect, useState } from 'react';
import { fetchRoleTypes } from 'api/roleType';

function useRoleTypes() {
  const [roleTypes, setRoleTypes] = useState([]);

  useEffect(() => {
    fetchRoleTypes()
      .then((roleTypes) => {
        setRoleTypes(roleTypes);
      });
  }, []);

  return [roleTypes, setRoleTypes];
}

export default useRoleTypes;
