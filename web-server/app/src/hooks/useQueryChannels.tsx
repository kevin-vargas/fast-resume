import { useQuery } from "react-query"
import Config from "../config"
import { Channels } from "../model"
const toJson = (res: Response) => res.json()
const queryChannels = () => fetch(Config.channelsUri, {credentials: 'include'}).then(toJson)
// TODO: generic querys
export default () => {
    const { data } = useQuery<Channels>(
        ['channels'],
        queryChannels,
        {staleTime: 5*60*1000}, 
    )
    return data
}