// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type { NextApiRequest, NextApiResponse } from 'next'
import externalApi from "../../app/externalApi";

type Data = {
    message: string
}

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
    if (req.method === 'POST') {

        const body = req.body

        const r = await externalApi.post('register', body)

        const data = r.data as Data

        res.status(r.status).json(data)
    }else {
        res.status(415).json({ message: '' } as Data)
    }
  
}
