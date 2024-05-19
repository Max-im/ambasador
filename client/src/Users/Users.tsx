import { Table, TableCell, TableContainer, TableFooter, TablePagination, TableRow } from '@mui/material';
import axios from 'axios';
import React, { useEffect, useState } from 'react';
import { Link } from 'react-router-dom';

export default function Users() {
  const [page, setPage] = useState(0);
  const [perPage, setPerPage] = useState(10);
  const [users, setUsers] = useState([]);
  
  useEffect(() => {
    axios.get('/ambassadors').then(({data}) => setUsers(data))
  }, []);

  return <div>Users:
      <TableContainer sx={{mt:2}}>
        <Table>
          <TableRow>
            <TableCell>#</TableCell>
            <TableCell>Name</TableCell>
            <TableCell>Email</TableCell>
            <TableCell>Action</TableCell>
          </TableRow>
          {users.slice(page*perPage, (page+1)*perPage).map((user: any) => (
            <TableRow key={user.id}>
              <TableCell>{user.id}</TableCell>
              <TableCell>{user.first_name} {user.last_name}</TableCell>
              <TableCell>{user.email}</TableCell>
              <TableCell>
                <Link to={`/users/${user.id}/links`}>View</Link>
              </TableCell>
            </TableRow>
          ))}
          <TableFooter>
            <TablePagination
              component="div"
              count={users.length}
              rowsPerPageOptions={[10, 20, 30]}
              onRowsPerPageChange={(e) => setPerPage(Number(e.target.value))}
              page={page}
              rowsPerPage={perPage}
              onPageChange={(_, page) => setPage(page)}
            />

          </TableFooter>
        </Table>
      </TableContainer>
  </div>;
}
