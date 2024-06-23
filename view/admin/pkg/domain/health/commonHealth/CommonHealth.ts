import type {CommonHealthType} from "~/pkg/domain/health/commonHealth/CommonHealthType";

export type CommonHealth = {
  "healthId": number;
  "name": string;
  "commonHealthType": CommonHealthType;
}
