import React, { useEffect, useState } from 'react'
import { useParams } from 'react-router-dom';
import axios from 'axios';
import { Table, TableCell, TableContainer, TableRow } from '@mui/material';
import { IOrder } from '../Order/Order';

interface ILinkItem {
    id: number;
    code: string;
    orders: IOrder[]
}

export default function UserLinks() {
    const { userId } = useParams()
    const [links, setLinks] = useState([])

console.log(userId)
    useEffect(() => {
        axios.get(`/user/${userId}/links`).then(({data}) => setLinks(data))
      }, []);

  return (
    <div>Links:
        <TableContainer sx={{mt:2}}>
        <Table>
          <TableRow>
            <TableCell>#</TableCell>
            <TableCell>Code</TableCell>
            <TableCell>Email</TableCell>
          </TableRow>
          {links.map((link: ILinkItem) => (
            <TableRow key={link.id}>
              <TableCell>{link.id}</TableCell>
              <TableCell>{link.code}</TableCell>
            </TableRow>
          ))}
          
        </Table>
      </TableContainer>
    </div>
  )
}
