using System;
using SQLite.Net;
using MvvmCross.Plugins.Sqlite;
using MvvmCross.Plugins.File;
using System.Collections.Generic;
using System.Linq;
using Newtonsoft.Json;

namespace PitWarner
{
    public class DataBaseService : IDataBaseService
    {
        SQLiteConnection _connection;
        IMvxFileStore _fileStore;
        IMvxSqliteConnectionFactory _sqliteConnectionFactory;

        public DataBaseService(IMvxSqliteConnectionFactory sqliteConnectionFactory, IMvxFileStore fileStore)
        {
            _sqliteConnectionFactory = sqliteConnectionFactory;
            _fileStore = fileStore;

            _connection = _sqliteConnectionFactory.GetConnection(_fileStore.NativePath(Variables.DataBaseName));
        }

        public void InitDB()
        {
            //_fileStore.DeleteFile(Variables.DataBaseName);
            //_fileStore.WriteFile (Variables.DataBaseName, string.Empty);
            //_connection.CreateTable<PitModel> ();
        }

        #region IDataBaseService implementation

        public void SaveData(System.Collections.Generic.IList<PitModel> pits)
        {
            //_connection.InsertAll (pits);
            var json = JsonConvert.SerializeObject (pits);
            _fileStore.WriteFile (_fileStore.NativePath(Variables.DataBaseName), json);
        }

        public System.Collections.Generic.List<PitModel> ReadData()
        {
            //return _connection.Table<PitModel> ().ToList();
            string json;
            _fileStore.TryReadTextFile (_fileStore.NativePath (Variables.DataBaseName), out json);
            var pits = JsonConvert.DeserializeObject<List<PitModel>> (json);
            return pits;
        }

        #endregion
    }
}

