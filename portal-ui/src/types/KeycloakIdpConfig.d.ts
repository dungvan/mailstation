export type KeycloakIdpHint = {
  providerId: string;
  displayName: string;
}

export type KeycloakIdpConfig = {
  issuer: string;
  clientId: string;
  clientSecret: string;
  redirectUri: string;
  scopes: string[];
  postLogoutRedirectUri: string;
  responseType?: string;
}
