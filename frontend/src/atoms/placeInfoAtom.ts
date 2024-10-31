import { atom } from "jotai";
import { GetPlaceInfoList } from "../../wailsjs/go/main/App";

const placeInfoAtom = atom(async (get) => {
	return await GetPlaceInfoList();
});

export default placeInfoAtom;
