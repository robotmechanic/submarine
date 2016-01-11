// This file was generated by typhen-api

using System;
using System.Collections.Generic;
using System.Linq;

namespace TyphenApi
{
    public interface ISerializationInfo
    {
        string PropertyName { get; }
        bool IsOptional { get; }
        System.Type ValueType { get; }

        object GetValue(object obj);
        void SetValue(object obj, object value);
    }

    public class SerializationInfo<ClassT, ValueT> : ISerializationInfo
    {
        readonly Func<ClassT, ValueT> getValueFunc;
        readonly Action<ClassT, ValueT> setValueFunc;

        public string PropertyName { get; private set; }
        public bool IsOptional { get; private set; }
        public System.Type ValueType { get; private set; }

        public SerializationInfo(string propertyName, bool isOptional,
            Func<ClassT, ValueT> getValueFunc, Action<ClassT, ValueT> setValueFunc)
        {
            this.getValueFunc = getValueFunc;
            this.setValueFunc = setValueFunc;
            PropertyName = propertyName;
            IsOptional = isOptional;
            ValueType = typeof(ValueT);
        }

        public object GetValue(object obj)
        {
            return getValueFunc((ClassT)obj);
        }

        public void SetValue(object obj, object rawValue)
        {
            ValueT value;

            try
            {
                value = (ValueT)rawValue;
            }
            catch (InvalidCastException)
            {
                value = (ValueT)Convert.ChangeType(rawValue, ValueType);
            }

            setValueFunc((ClassT)obj, value);
        }
    }

    public class SerializationInfoForList<ClassT, ValueT> : ISerializationInfo
    {
        readonly Func<ClassT, List<ValueT>> getValueFunc;
        readonly Action<ClassT, List<ValueT>> setValueFunc;

        public string PropertyName { get; private set; }
        public bool IsOptional { get; private set; }
        public System.Type ValueType { get; private set; }

        public SerializationInfoForList(string propertyName, bool isOptional,
            Func<ClassT, List<ValueT>> getValueFunc, Action<ClassT, List<ValueT>> setValueFunc)
        {
            this.getValueFunc = getValueFunc;
            this.setValueFunc = setValueFunc;
            PropertyName = propertyName;
            IsOptional = isOptional;
            ValueType = typeof(List<ValueT>);
        }

        public object GetValue(object obj)
        {
            return getValueFunc((ClassT)obj).OfType<object>();
        }

        public void SetValue(object obj, object rawValue)
        {
            var value = ((IEnumerable<object>)rawValue).OfType<ValueT>().ToList();
            setValueFunc((ClassT)obj, value);
        }
    }
}
