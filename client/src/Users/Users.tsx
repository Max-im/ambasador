import { Table, TableCell, TableContainer, TableRow } from '@mui/material';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

export default function Users() {
  const [users, setUsers] = useState([]);
  useEffect(() => {
    axios.get('/ambassadors').then(({data}) => setUsers(data))
  }, []);

  return <div>Users:
      <TableContainer sx={{mt:2}}>
        <Table>
          <TableRow>
            <TableCell>First Name</TableCell>
            <TableCell>Last Name</TableCell>
            <TableCell>Email</TableCell>
            <TableCell>Action</TableCell>
          </TableRow>
          {users.map((user: any) => (
            <TableRow key={user.id}>
              <TableCell>{user.first_name}</TableCell>
              <TableCell>{user.last_name}</TableCell>
              <TableCell>{user.email}</TableCell>
              <TableCell>
                <Link to={`/users/${user.id}`}>View</Link>
              </TableCell>
            </TableRow>
          ))}
        </Table>
      </TableContainer>
  </div>;
}
