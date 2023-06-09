import { type TrustLevel } from "../types/dto/trust_level";
import { baseApiClient } from "./client"

export const fetchTrustLevels = async () => {
  const result = await baseApiClient.get("/trust");
  const data: any[] = result.data;
  const trust_levels: TrustLevel[] = data.map((value) => {
    return {
      extrasenseId: value["extrasense_id"],
      trust: value["trust"],
    };
  });
  return trust_levels;
}
