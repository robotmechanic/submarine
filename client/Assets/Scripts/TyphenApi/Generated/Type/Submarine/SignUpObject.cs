// This file was generated by typhen-api

using System.Collections.Generic;

namespace TyphenApi.Type.Submarine
{
    public partial class SignUpObject : TyphenApi.TypeBase<SignUpObject>
    {
        protected static readonly SerializationInfo<SignUpObject, TyphenApi.Type.Submarine.LoggedInUser> user = new SerializationInfo<SignUpObject, TyphenApi.Type.Submarine.LoggedInUser>("user", false, (x) => x.User, (x, v) => x.User = v);
        public TyphenApi.Type.Submarine.LoggedInUser User { get; set; }
        protected static readonly SerializationInfo<SignUpObject, string> authToken = new SerializationInfo<SignUpObject, string>("auth_token", false, (x) => x.AuthToken, (x, v) => x.AuthToken = v);
        public string AuthToken { get; set; }
    }
}
