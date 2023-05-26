import { TrustLevel } from "../types/dto/trust_level";
import { baseApiClient } from "./client"

export const fetchTrustLevels = async () => {
  let result = await baseApiClient.get("/trust");
  let data: any[] = result.data;
  let trust_levels: TrustLevel[] = data.map((value) => {
    return {
      extrasenseId: value["extrasense_id"],
      trust: value["trust"],
    };
  });
  sessionStorage.setItem("trust-levels", JSON.stringify(trust_levels))
}
